package math

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
	"math"
)

// 求余指令
type DREM struct {
	base.NoOperandsInstruction
}
type FREM struct {
	base.NoOperandsInstruction
}
type IREM struct {
	base.NoOperandsInstruction
}
type LREM struct {
	base.NoOperandsInstruction
}

func (this *IREM) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

func (this *LREM) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException:/by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

// 对于双精度浮点类型，应该使用mod函数，因为由Infinity无穷大值，所以/0也不会异常抛出
func (this *DREM) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
func (this *FREM) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := math.Mod(float64(v1), float64(v2))
	stack.PushFloat(float32(result))
}
