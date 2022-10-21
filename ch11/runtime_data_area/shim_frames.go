package runtime_data_area

import "go-jvm/ch11/runtime_data_area/heap"

func NewShimFrame(thread *Thread, ops *OperandStack) *Frame {
	return &Frame{
		thread:       thread,
		method:       heap.ShimReturnMethod(),
		operandStack: ops,
	}
}
