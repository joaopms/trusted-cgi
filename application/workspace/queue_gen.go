package workspace

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Task) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "environment":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Environment")
				return
			}
			if z.Environment == nil {
				z.Environment = make(map[string]string, zb0002)
			} else if len(z.Environment) > 0 {
				for key := range z.Environment {
					delete(z.Environment, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Environment")
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Environment", za0001)
					return
				}
				z.Environment[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Task) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "environment"
	err = en.Append(0x81, 0xab, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Environment)))
	if err != nil {
		err = msgp.WrapError(err, "Environment")
		return
	}
	for za0001, za0002 := range z.Environment {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Environment")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Environment", za0001)
			return
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Task) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.Environment != nil {
		for za0001, za0002 := range z.Environment {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	return
}
