package constants

import "go-jvm/ch10/instructions/base"
import "go-jvm/ch10/runtime_data_area"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (this *NOP) Execute(frame *runtime_data_area.Frame) {
	// really do nothing
}
