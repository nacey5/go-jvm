package control

import "go-jvm/ch08-2/instructions/base"
import "go-jvm/ch08-2/runtime_data_area"

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (this *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.low = reader.ReadInt32()
	this.high = reader.ReadInt32()
	jumpOffsetsCount := this.high - this.low + 1
	this.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (this *TABLE_SWITCH) Execute(frame *runtime_data_area.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= this.low && index <= this.high {
		offset = int(this.jumpOffsets[index-this.low])
	} else {
		offset = int(this.defaultOffset)
	}

	base.Branch(frame, offset)
}
