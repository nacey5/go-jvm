package runtime_data_area

// num存放整数
// ref存放引用的地址
type Slot struct {
	num int32
	ref *Object
}
