package classfile

import (
	"go-jvm/ch03/util"
	"math"
)

// 常量池信息结构体
type ConstantClassInfo struct {
	constantPool ConstantPool
	nameIndex    uint16
}

type ConstantDoubleInfo struct {
	val float64
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantIntegerInfo struct {
	val int32
}

type ConstantLongInfo struct {
	val int64
}

// 字段名称和描述符
// B/S/C/I/J/F/D  -------byte,short,char,int,long,float,double
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

type ConstantStringInfo struct {
	constantPool ConstantPool
	stringIndex  uint16
}

type ConstantUtf8Info struct {
	str string
}

// 对于ConstantIntegerInfo来讲，读取一个int32位的数据，对于boolean，byte,short,char,也可以使用32位进行存储
func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = int32(bytes)
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = math.Float32frombits(bytes)
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = int64(bytes)
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = math.Float64frombits(bytes)
}

func (this *ConstantUtf8Info) readInfo(reader *ClassReader) {
	//第一个16位表示长度
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	this.str = util.DecodeMUTF8(bytes)
}

func (this *ConstantStringInfo) readInfo(reader *ClassReader) {
	this.stringIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (this *ConstantStringInfo) String() string {
	return this.constantPool.getUtf8(this.stringIndex)
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
}

func (this *ConstantClassInfo) Name() string {
	return this.constantPool.getUtf8(this.nameIndex)
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
	this.descriptorIndex = reader.readUint16()
}
