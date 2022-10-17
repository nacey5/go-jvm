package runtime_data_area

import "go-jvm/ch08-2/runtime_data_area/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
