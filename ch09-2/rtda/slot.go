package rtda

import "go-jvm/ch09-2/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
