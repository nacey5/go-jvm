package classfile

import "encoding/binary"

// u1,u2,u4,三种数据类型代表了三种数据类型表示1，2，4无符号整数
type ClassReader struct {
	data []byte
}

// u1读取u1类型数据
func (this *ClassReader) readUint8() uint8 {
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

// u2 使用go特性reslice语法跳过已经读过的数据，读取u2类型数据
func (this *ClassReader) readUint16() uint16 {
	//读取一个Uint16类型的数据出来
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

// 粗去uint16表，大小由第一个数字决定
func (this *ClassReader) readUint16s() []uint16 {
	count := this.readUint16()
	//初始化
	ms := make([]uint16, count)
	for i := range ms {
		ms[i] = this.readUint16()
	}
	return ms
}

func (this *ClassReader) readBytes(length uint32) []byte {
	bytes := this.data[:length]
	this.data = this.data[length:]
	return bytes
}
