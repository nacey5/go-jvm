package heap

// 几个基本数据类型的
// 这8种分别对应8条stroe和load指令
// 对于boolean类型通过bytes就可以实现
func (this *Object) Bytes() []int8 {
	return this.data.([]int8)
}

func (this *Object) Shorts() []int16 {
	return this.data.([]int16)
}
func (this *Object) Ints() []int32 {
	return this.data.([]int32)
}
func (this *Object) Longs() []int64 {
	return this.data.([]int64)
}
func (this *Object) Chars() []uint16 {
	return this.data.([]uint16)
}
func (this *Object) Floats() []float32 {
	return this.data.([]float32)
}
func (this *Object) Doubles() []float64 {
	return this.data.([]float64)
}
func (this *Object) Refs() []*Object {
	return this.data.([]*Object)
}

// 基本数据类型数组的长度
func (this *Object) ArrayLength() int32 {
	switch this.data.(type) {
	case []int8:
		return int32(len(this.data.([]int8)))
	case []int16:
		return int32(len(this.data.([]int16)))
	case []int32:
		return int32(len(this.data.([]int32)))
	case []int64:
		return int32(len(this.data.([]int64)))
	case []uint16:
		return int32(len(this.data.([]uint16)))
	case []float32:
		return int32(len(this.data.([]float32)))
	case []float64:
		return int32(len(this.data.([]float64)))
	case []*Object:
		return int32(len(this.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
