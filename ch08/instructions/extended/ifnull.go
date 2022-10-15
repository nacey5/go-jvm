package extended

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

// Branch is reference is null
type IFNULL struct {
	base.BranchInstruction
}

// Branch is reference is not null
type IFNONNULL struct {
	base.BranchInstruction
}

// ifnull 和ifnonnull指令把栈顶弹出
func (this *IFNULL) Execute(frame *runtime_data_area.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFNONNULL) Execute(frame *runtime_data_area.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, this.Offset)
	}
}
