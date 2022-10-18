package constants

import "go-jvm/ch09/instructions/base"
import "go-jvm/ch09/runtime_data_area"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (this *NOP) Execute(frame *runtime_data_area.Frame) {
	// really do nothing
}
