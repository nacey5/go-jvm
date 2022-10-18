package lang

import (
	"go-jvm/ch09/native"
	"go-jvm/ch09/runtime_data_area"
	"math"
)

func init() {
	native.Registry("java/lang/Float", "floatToRowIntBits", "(F)I", floatToRowIntBits)
}

// public static native int floatToRowIntBits(float value)
func floatToRowIntBits(frame *runtime_data_area.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}
