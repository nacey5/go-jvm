package constants

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (this *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
