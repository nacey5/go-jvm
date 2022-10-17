package extended

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (this *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	this.offset = int(reader.ReadInt32())
}
func (this *GOTO_W) Execute(frame *runtime_data_area.Frame) {
	base.Branch(frame, this.offset)
}
