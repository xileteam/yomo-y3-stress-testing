package decoder

import (
	"fmt"
)

const MSB byte = 0x80

// DropMSB 描述了`0111 1111`, 用于去除标识位使用
const DropMSB = 0x3F

// DropMSBArrayFlag 描述了`0011 1111`, 用于去除标识位使用
const DropMSBArrayFlag = 0x3F

// ArrayFlag 描述了`0100 0000`, 用于表示该节点的Value为Slice类型
const ArrayFlag = 0x40

// Tag represents the Tag of TLV
// MSB used to represent the packet type, 0x80 means a node packet, otherwise is a primitive packet
// Low 7 bits represent Sequence ID, like `key` in JSON format
type Tag struct {
	raw byte
}

// IsNode returns true is MSB is 1.
func (t *Tag) IsNode() bool {
	return t.raw&MSB == MSB
}

// SeqID get the sequence ID, as key in JSON format
func (t *Tag) SeqID() byte {
	//return t.raw & utils.DropMSB
	return t.raw & DropMSBArrayFlag
}

func (t *Tag) String() string {
	return fmt.Sprintf("Tag: raw=%4b, SeqID=%v", t.raw, t.SeqID())
}

// NewTag create a NodePacket Tag field
func NewTag(b byte) *Tag {
	return &Tag{raw: b}
}

func (t *Tag) IsArray() bool {
	return t.raw&ArrayFlag == ArrayFlag
}
