package heap

import "go-jvm/ch09/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}
func (this *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		this.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (this *Field) IsVolatile() bool {
	return 0 != this.accessFlags&ACC_VOLATILE
}
func (this *Field) IsTransient() bool {
	return 0 != this.accessFlags&ACC_TRANSIENT
}
func (this *Field) IsEnum() bool {
	return 0 != this.accessFlags&ACC_ENUM
}

func (this *Field) ConstValueIndex() uint {
	return this.constValueIndex
}
func (this *Field) SlotId() uint {
	return this.slotId
}
func (this *Field) isLongOrDouble() bool {
	return this.descriptor == "J" || this.descriptor == "D"
}
