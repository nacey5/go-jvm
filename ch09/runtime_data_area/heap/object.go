package heap

type Object struct {
	class *Class
	data  interface{} // Slots for Object, []int32 for int[] ...
	extra interface{}
}

func (this *Object) Extra() interface{} {
	return this.extra
}

func (this *Object) SetExtra(extra interface{}) {
	this.extra = extra
}

// create normal (non-array) object
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

func (this *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(this.class)
}

// reflection
func (this *Object) GetRefVar(name, descriptor string) *Object {
	field := this.class.getField(name, descriptor, false)
	slots := this.data.(Slots)
	return slots.GetRef(field.slotId)
}
func (this *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := this.class.getField(name, descriptor, false)
	slots := this.data.(Slots)
	slots.SetRef(field.slotId, ref)
}
