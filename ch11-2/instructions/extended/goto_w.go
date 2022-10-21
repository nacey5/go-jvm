package extended

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"

// Branch always (wide index)
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *runtime_data_area.Frame) {
	base.Branch(frame, self.offset)
}
