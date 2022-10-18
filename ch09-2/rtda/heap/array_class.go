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
		return &Object{this, make([]int8, count), nil}
	case "[B":
		return &Object{this, make([]int8, count), nil}
	case "[C":
		return &Object{this, make([]uint16, count), nil}
	case "[S":
		return &Object{this, make([]int16, count), nil}
	case "[I":
		return &Object{this, make([]int32, count), nil}
	case "[J":
		return &Object{this, make([]int64, count), nil}
	case "[F":
		return &Object{this, make([]float32, count), nil}
	case "[D":
		return &Object{this, make([]float64, count), nil}
	default:
		return &Object{this, make([]*Object, count), nil}
	}
}
