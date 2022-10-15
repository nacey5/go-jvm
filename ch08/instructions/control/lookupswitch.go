package control

import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

// 这个也需要先跳过padding
func (this *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffsets = reader.ReadInt32s(this.npairs * 2)
}

// 有点像Map，key是case值，value是跳转偏移量
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
