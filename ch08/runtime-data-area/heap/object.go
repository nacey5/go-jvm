package heap

// class 存放对象的class指针，一个存放对象实例
// []slot --> data
type Object struct {
	class *Class
	data  interface{}
}

// 判断类型
func (this Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(this.class)
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

// getters
func (this *Object) Class() *Class {
	return this.class
}
func (this *Object) Fields() Slots {
	return this.data.(Slots)
}
