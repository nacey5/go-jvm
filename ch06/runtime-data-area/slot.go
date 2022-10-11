package runtime_data_area

import "go-jvm/ch06/runtime-data-area/heap"

// num存放整数
// ref存放引用的地址
type Slot struct {
	num int32
	ref *heap.Object
}
