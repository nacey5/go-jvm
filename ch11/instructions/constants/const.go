package constants

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Push null
type ACONST_NULL struct{ base.NoOperandsInstruction }

func (this *ACONST_NULL) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Push double
type DCONST_0 struct{ base.NoOperandsInstruction }

func (this *DCONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperandsInstruction }

func (this *DCONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct{ base.NoOperandsInstruction }

func (this *FCONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperandsInstruction }

func (this *FCONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperandsInstruction }

func (this *FCONST_2) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant
type ICONST_M1 struct{ base.NoOperandsInstruction }

func (this *ICONST_M1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperandsInstruction }

func (this *ICONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperandsInstruction }

func (this *ICONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperandsInstruction }

func (this *ICONST_2) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperandsInstruction }

func (this *ICONST_3) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperandsInstruction }

func (this *ICONST_4) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperandsInstruction }

func (this *ICONST_5) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct{ base.NoOperandsInstruction }

func (this *LCONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperandsInstruction }

func (this *LCONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushLong(1)
}
