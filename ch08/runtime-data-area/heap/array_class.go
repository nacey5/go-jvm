package heap

func (this *Class) NewArray(count uint) *Object {
	//如果不是数组
	if !this.IsArray() {
		panic("Not array class:" + this.name)
	}
	switch this.Name() {
	case "[Z":
		return &Object{this, make([]int8, count)}
	case "[B":
		return &Object{this, make([]int8, count)}
	case "[C":
		return &Object{this, make([]uint16, count)}
	case "[S":
		return &Object{this, make([]int16, count)}
	case "[I":
		return &Object{this, make([]int32, count)}
	case "[J":
		return &Object{this, make([]int64, count)}
	case "[F":
		return &Object{this, make([]float32, count)}
	case "[D":
		return &Object{this, make([]float64, count)}
	default:
		return &Object{this, make([]*Object, count)}
	}
}
