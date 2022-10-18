package extended

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (this *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	this.offset = int(reader.ReadInt32())
}
func (this *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.offset)
}
