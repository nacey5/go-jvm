package constants

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}
type DCONST_0 struct {
	base.NoOperandsInstruction
}
type DCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_0 struct {
	base.NoOperandsInstruction
}
type FCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_M1 struct {
	base.NoOperandsInstruction
}
type ICONST_0 struct {
	base.NoOperandsInstruction
}
type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (this *ACONST_NULL) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (this *DCONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (this *ICONST_M1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(-1)
}
func (this *DCONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushDouble(1.0)
}
func (this *FCONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushFloat(0.0)
}
func (this *FCONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushFloat(1.0)
}
func (this *FCONST_2) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushFloat(2.0)
}
func (this *ICONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(0)
}
func (this *ICONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(1)
}
func (this *ICONST_2) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(2)
}
func (this *ICONST_3) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(3)
}
func (this *ICONST_4) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(4)
}
func (this *ICONST_5) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushInt(5)
}
func (this *LCONST_0) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushLong(0)
}
func (this *LCONST_1) Execute(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushLong(1)
}
