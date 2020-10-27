package mapping

import (
	"github.com/10cella/yomo-y3-stress-testing/internal/utils"
	"reflect"

	y3 "github.com/yomorun/yomo-codec-golang"
)

// temp: test...
func NewDataCodec(data map[byte]interface{}) []byte {
	npn := EncodeMap(data, nil)
	return npn.Encode()
}

// temp: test...
func EncodeMap(data map[byte]interface{}, wrapper *y3.NodePacketEncoder) *y3.NodePacketEncoder {
	if wrapper == nil {
		wrapper = y3.NewNodePacketEncoder(0x01)
	}

	for k, v := range data {
		var ppe = y3.NewPrimitivePacketEncoder(int(k))
		//fmt.Printf("#99 EncodeMap reflect.TypeOf(v)=%v\n", reflect.TypeOf(v))
		switch value := v.(type) {
		case string:
			ppe.SetStringValue(value)
			wrapper.AddPrimitivePacket(ppe)
		case int32:
			ppe.SetInt32Value(v.(int32))
			wrapper.AddPrimitivePacket(ppe)
		case uint32:
			ppe.SetUInt32Value(v.(uint32))
			wrapper.AddPrimitivePacket(ppe)
		case int64:
			ppe.SetInt64Value(v.(int64))
			wrapper.AddPrimitivePacket(ppe)
		case uint64:
			ppe.SetUInt64Value(v.(uint64))
			wrapper.AddPrimitivePacket(ppe)
		case float32:
			ppe.SetFloat32Value(v.(float32))
			wrapper.AddPrimitivePacket(ppe)
		case float64:
			ppe.SetFloat64Value(v.(float64))
			wrapper.AddPrimitivePacket(ppe)
		case []interface{}:
			var root = y3.NewNodeArrayPacketEncoder(int(k))
			for _, item := range v.([]interface{}) {
				switch reflect.TypeOf(item).Kind() {
				case reflect.Map:
					handleMapArray(root, item)
				}
			}
			wrapper.AddNodePacket(root)

		case map[byte][]interface{}:
			//fmt.Printf("map[byte][]interface{}=%v\n", v)
			var root = y3.NewNodePacketEncoder(int(k))
			for ku, uv := range v.(map[byte][]interface{}) {
				w := y3.NewNodeArrayPacketEncoder(int(ku))
				for _, kk := range uv {
					var ppe = y3.NewPrimitivePacketEncoder(0x00)
					switch reflect.ValueOf(kk).Kind() {
					case reflect.String:
						ppe.SetStringValue(kk.(string))
					case reflect.Int32:
						ppe.SetInt32Value(kk.(int32))
					case reflect.Uint32:
						ppe.SetUInt32Value(kk.(uint32))
					case reflect.Int64:
						ppe.SetInt64Value(kk.(int64))
					case reflect.Uint64:
						ppe.SetUInt64Value(kk.(uint64))
					case reflect.Float32:
						ppe.SetFloat32Value(kk.(float32))
					case reflect.Float64:
						ppe.SetFloat64Value(kk.(float64))
					}
					w.AddPrimitivePacket(ppe)
				}
				root.AddNodePacket(w)
			}
			wrapper.AddNodePacket(root)
		case map[byte]interface{}:
			var root = y3.NewNodePacketEncoder(int(k))
			for kk, vv := range v.(map[byte]interface{}) {
				var ppe = y3.NewPrimitivePacketEncoder(int(kk))
				switch reflect.ValueOf(vv).Kind() {
				case reflect.String:
					ppe.SetStringValue(vv.(string))
				case reflect.Int32:
					ppe.SetInt32Value(vv.(int32))
				case reflect.Uint32:
					ppe.SetUInt32Value(vv.(uint32))
				case reflect.Int64:
					ppe.SetInt64Value(vv.(int64))
				case reflect.Uint64:
					ppe.SetUInt64Value(vv.(uint64))
				case reflect.Float32:
					ppe.SetFloat32Value(vv.(float32))
				case reflect.Float64:
					ppe.SetFloat64Value(vv.(float64))
				}
				root.AddPrimitivePacket(ppe)
			}
			wrapper.AddNodePacket(root)
		}
	}

	return wrapper
}

func handleMapArray(node *y3.NodePacketEncoder, item interface{}) {
	//var nn = y3.NewNodeArrayPacketEncoder(0x00)
	if m, ok := item.(map[byte]string); ok {
		setStringToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]int32); ok {
		setInt32ToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]uint32); ok {
		setUInt32ToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]int64); ok {
		setInt64ToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]uint64); ok {
		setUInt64ToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]float32); ok {
		setFloat32ToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]float64); ok {
		setFloat64ToArray(node, m)
		return
	}
	if m, ok := item.(map[byte]interface{}); ok {
		setInterfaceToArray(node, m)
		return
	}
	//root.AddNodePacket(nn)
}

func setStringToArray(node *y3.NodePacketEncoder, m map[byte]string) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetStringValue(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setInt32ToArray(node *y3.NodePacketEncoder, m map[byte]int32) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetInt32Value(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setUInt32ToArray(node *y3.NodePacketEncoder, m map[byte]uint32) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetUInt32Value(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setInt64ToArray(node *y3.NodePacketEncoder, m map[byte]int64) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetInt64Value(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setUInt64ToArray(node *y3.NodePacketEncoder, m map[byte]uint64) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetUInt64Value(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setFloat32ToArray(node *y3.NodePacketEncoder, m map[byte]float32) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetFloat32Value(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setFloat64ToArray(node *y3.NodePacketEncoder, m map[byte]float64) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		var n1 = y3.NewPrimitivePacketEncoder(int(kk))
		n1.SetFloat64Value(vv)
		mm.AddPrimitivePacket(n1)
	}
	node.AddNodePacket(mm)
}

func setInterfaceToArray(node *y3.NodePacketEncoder, m map[byte]interface{}) {
	var mm = y3.NewNodePacketEncoder(utils.KeyOfArrayItem)
	for kk, vv := range m {
		switch reflect.TypeOf(vv).Kind() {
		case reflect.String:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetStringValue(vv.(string))
			mm.AddPrimitivePacket(n1)
		case reflect.Int32:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetInt32Value(vv.(int32))
			mm.AddPrimitivePacket(n1)
		case reflect.Uint32:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetUInt32Value(vv.(uint32))
			mm.AddPrimitivePacket(n1)
		case reflect.Int64:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetInt64Value(vv.(int64))
			mm.AddPrimitivePacket(n1)
		case reflect.Uint64:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetUInt64Value(vv.(uint64))
			mm.AddPrimitivePacket(n1)
		case reflect.Float32:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetFloat32Value(vv.(float32))
			mm.AddPrimitivePacket(n1)
		case reflect.Float64:
			var n1 = y3.NewPrimitivePacketEncoder(int(kk))
			n1.SetFloat64Value(vv.(float64))
			mm.AddPrimitivePacket(n1)
		}
		//fmt.Printf("#200 kk=%#x, vv=%v\n", kk, vv)
		//fmt.Printf("#200 reflect.TypeOf(item).Kind()=%v\n", reflect.TypeOf(vv).Kind())
	}
	node.AddNodePacket(mm)
}
