package comparisons

import (
	"go-jvm/ch07/instructions/base"
	runtime_data_area "go-jvm/ch07/runtime-data-area"
)

// 浮点数可能产生NAN，所以可能无法比较
type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *runtime_data_area.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

// 有NAN时，fcmp指令比较的结果是1
func (this *FCMPG) Execute(frame *runtime_data_area.Frame) {
	_fcmp(frame, true)
}

// 有NAN时，fcmp指令比较的结果是-1
func (this *FCMPL) Execute(frame *runtime_data_area.Frame) {
	_fcmp(frame, false)
}
