package runtime_data_area

import "go-jvm/ch09/runtime_data_area/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
