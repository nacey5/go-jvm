package base

type BytecodeReader struct {
	//code字段存放字节码
	code []byte
	//pc记录读取到哪个字节
	pc int
}

func (this *BytecodeReader) Code() []byte {
	return this.code
}

func (this *BytecodeReader) SetCode(code []byte) {
	this.code = code
}

func (this *BytecodeReader) Pc() int {
	return this.pc
}

func (this *BytecodeReader) SetPc(pc int) {
	this.pc = pc
}

// Reset 避免每次访问都创建一个实例--单例模式
func (this *BytecodeReader) Reset(code []byte, pc int) {
	this.code = code
	this.pc = pc
}

// 实现一系列Read()方法
func (this *BytecodeReader) ReadUint8() uint8 {
	i := this.code[this.pc]
	this.pc++
	return i
}
func (this *BytecodeReader) ReadInt8() int8 {
	return int8(this.ReadUint8())
}
func (this *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(this.ReadUint8())
	byte2 := uint16(this.ReadUint8())
	return (byte1 << 8) | byte2
}
func (this *BytecodeReader) ReadInt16() int16 {
	return int16(this.ReadUint16())
}
func (this *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(this.ReadUint8())
	byte2 := int32(this.ReadUint8())
	byte3 := int32(this.ReadUint8())
	byte4 := int32(this.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// 只有tableswitch和lookupswitch需要使用
// 执行跳转所需的字节码偏移量
func (this *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = this.ReadInt32()
	}
	return ints
}

// 只有tableswitch和lookupswitch需要使用
// tableswitch后面有0~3字节的padding，保证defaultOffset的地址是4的倍数
func (this *BytecodeReader) SkipPadding() {
	for this.pc%4 != 0 {
		this.ReadUint8()
	}
}
