package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// u1
func (this *ClassReader) readUint8() uint8 {
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

// u2
func (this *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(this.data)
	this.data = this.data[2:]
	return val
}

// u4
func (this *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(this.data)
	this.data = this.data[4:]
	return val
}

func (this *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(this.data)
	this.data = this.data[8:]
	return val
}

func (this *ClassReader) readUint16s() []uint16 {
	n := this.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = this.readUint16()
	}
	return s
}

func (this *ClassReader) readBytes(n uint32) []byte {
	bytes := this.data[:n]
	this.data = this.data[n:]
	return bytes
}
