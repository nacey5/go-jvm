package control

import "go-jvm/ch10/instructions/base"
import "go-jvm/ch10/runtime_data_area"

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (this *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffsets = reader.ReadInt32s(this.npairs * 2)
}

func (this *LOOKUP_SWITCH) Execute(frame *runtime_data_area.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < this.npairs*2; i += 2 {
		if this.matchOffsets[i] == key {
			offset := this.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(this.defaultOffset))
}
