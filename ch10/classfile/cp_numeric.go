package classfile

import "math"

/*
	CONSTANT_Integer_info {
	    u1 tag;
	    u4 bytes;
	}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = int32(bytes)
}
func (this *ConstantIntegerInfo) Value() int32 {
	return this.val
}

/*
	CONSTANT_Float_info {
	    u1 tag;
	    u4 bytes;
	}
*/
type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = math.Float32frombits(bytes)
}
func (this *ConstantFloatInfo) Value() float32 {
	return this.val
}

/*
	CONSTANT_Long_info {
	    u1 tag;
	    u4 high_bytes;
	    u4 low_bytes;
	}
*/
type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = int64(bytes)
}
func (this *ConstantLongInfo) Value() int64 {
	return this.val
}

/*
	CONSTANT_Double_info {
	    u1 tag;
	    u4 high_bytes;
	    u4 low_bytes;
	}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = math.Float64frombits(bytes)
}
func (this *ConstantDoubleInfo) Value() float64 {
	return this.val
}
