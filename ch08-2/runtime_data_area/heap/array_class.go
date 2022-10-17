package heap

func (this *Class) IsArray() bool {
	return this.name[0] == '['
}

func (this *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(this.name)
	return this.loader.LoadClass(componentClassName)
}

func (this *Class) NewArray(count uint) *Object {
	if !this.IsArray() {
		panic("Not array class: " + this.name)
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
