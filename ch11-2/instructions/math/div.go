package math

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"

// Divide double
type DDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type FDIV struct{ base.NoOperandsInstruction }

func (self *FDIV) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type IDIV struct{ base.NoOperandsInstruction }

func (self *IDIV) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

// Divide long
type LDIV struct{ base.NoOperandsInstruction }

func (self *LDIV) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
