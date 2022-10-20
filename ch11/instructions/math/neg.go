package math

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Negate double
type DNEG struct{ base.NoOperandsInstruction }

func (this *DNEG) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperandsInstruction }

func (this *FNEG) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperandsInstruction }

func (this *INEG) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperandsInstruction }

func (this *LNEG) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
