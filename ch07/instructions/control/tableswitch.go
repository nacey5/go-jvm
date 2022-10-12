package control

import (
	"go-jvm/ch07/instructions/base"
	runtime_data_area "go-jvm/ch07/runtime-data-area"
)

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
	jumpOffsetCount := this.high - this.low + 1
	this.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}

// 弹出一个int变量，然后看它是否在low和high给定的范围之内
// 如果在，从jumpOffset表中查出偏移量进行跳转，否则按照defaultOffset跳转
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
