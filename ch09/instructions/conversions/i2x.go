package conversions

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Convert int to byte
type I2B struct{ base.NoOperandsInstruction }

func (this *I2B) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	b := int32(int8(i))
	stack.PushInt(b)
}

// Convert int to char
type I2C struct{ base.NoOperandsInstruction }

func (this *I2C) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	c := int32(uint16(i))
	stack.PushInt(c)
}

// Convert int to short
type I2S struct{ base.NoOperandsInstruction }

func (this *I2S) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	s := int32(int16(i))
	stack.PushInt(s)
}

// Convert int to long
type I2L struct{ base.NoOperandsInstruction }

func (this *I2L) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	l := int64(i)
	stack.PushLong(l)
}

// Convert int to float
type I2F struct{ base.NoOperandsInstruction }

func (this *I2F) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	f := float32(i)
	stack.PushFloat(f)
}

// Convert int to double
type I2D struct{ base.NoOperandsInstruction }

func (this *I2D) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	d := float64(i)
	stack.PushDouble(d)
}
