package extended

import (
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

// 和goto的区别就是变成了四个字节
type GOTO_W struct {
	offset int
}

func (this *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	this.offset = int(reader.ReadInt32())
}

func (this *GOTO_W) Execute(frame *runtime_data_area.Frame) {
	base.Branch(frame, this.offset)
}
