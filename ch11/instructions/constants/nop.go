package constants

import "go-jvm/ch11/instructions/base"
import "go-jvm/ch11/runtime_data_area"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (this *NOP) Execute(frame *runtime_data_area.Frame) {
	// really do nothing
}
