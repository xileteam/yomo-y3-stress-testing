package decoder

import (
	"fmt"
	y3 "github.com/yomorun/yomo-codec-golang"
	"github.com/yomorun/yomo-codec-golang/pkg/codes"
	"github.com/yomorun/yomo-codec-golang/pkg/spec/encoding"
)

func TakeValueFromCodec(key byte, buf []byte) interface{} {
	//fmt.Printf("#101 buf %s\n", codes.FormatBytes(buf))
	var (
		Tag       *Tag   = nil
		Sbuf             = make([]byte, 0)
		Size      int32  = 0
		Length    int32  = 0
		LengthBuf []byte = make([]byte, 0)
	)

	for _, c := range buf {
		// tag
		if Tag == nil {
			Tag = NewTag(c)
			Sbuf = append(Sbuf, c)
			continue
		}

		// length
		if Size == 0 {
			LengthBuf = append(LengthBuf, c)
			Sbuf = append(Sbuf, c)
			length, size, err := DecodeLength(LengthBuf)
			if err != nil {
				continue
			}
			Length = length
			Size = size
			continue
		}

		//debug:
		//if Tag.SeqID() == 0x01 || Tag.SeqID() == 0x02 {
		//	fmt.Printf("#102 Tag.SeqID()=%#x, len(sbuf)=%v\n", Tag.SeqID(), len(sbuf))
		//}


		if key != Tag.SeqID() {
			//fmt.Printf("#101 Tag.SeqID()=%#x\n", Tag.SeqID())
			//如果key不匹配，则拦截后续的字节流进一步分析
			var newBuf []byte
			if len(buf) > int(1+Size+Length) {
				newBuf = buf[1+Size+Length:]
			} else {
				newBuf = buf[1+Size:]
			}

			return TakeValueFromCodec(key, newBuf)
		}

		Sbuf = append(Sbuf, c)

		// buf end, then handle sbuf
		if key == Tag.SeqID() && int32(len(Sbuf)) == 1+Size+Length {
			if Tag.IsNode() {
				// Decode Packet from sbuf
				packet, _, err := y3.DecodeNodePacket(Sbuf)
				if err != nil {
					//panic(errors.New("not a NodePacket"))
					fmt.Printf("#101.1 Tag.SeqID()=%#x, sbuf=%s error:%v\n",
						Tag.SeqID(), codes.FormatBytes(Sbuf), err)
					continue
				}

				//matching
				if ok, _, p := MatchingKey(key, packet); ok {
					return p
				}
			} else {
				packet, _, _, err := y3.DecodePrimitivePacket(Sbuf)
				if err != nil {
					fmt.Printf("#101.1 Tag.SeqID()=%#x, sbuf=%s error:%v\n",
						Tag.SeqID(), codes.FormatBytes(Sbuf), err)
					continue
				}
				return *packet
			}
		}
	}

	return nil
}

func DecodeLength(buf []byte) (length int32, size int32, err error) {
	varCodec := encoding.VarCodec{}
	err = varCodec.DecodePVarInt32(buf, &length)
	size = int32(varCodec.Size)
	return
}

func MatchingKey(key byte, node *y3.NodePacket) (flag bool, isNode bool, packet interface{}) {
	if len(node.PrimitivePackets) > 0 {
		for _, p := range node.PrimitivePackets {
			if key == p.SeqID() {
				return true, false, p
			}
		}
	}

	if len(node.NodePackets) > 0 {
		for _, n := range node.NodePackets {
			if key == n.SeqID() {
				return true, true, n
			}
			//return MatchingKey(key, &n)
			flag, isNode, packet = MatchingKey(key, &n)
			if flag {
				return
			}
		}
	}

	return false, false, nil
}
