package heap

import "go-jvm/ch06/classfile"

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

type Method struct {
	ClassMember
	maxStack  uint16
	maxLocals uint16
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (this *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttibute := cfMethod.CodeAttribute(); codeAttibute != nil {
		this.maxStack = codeAttibute.MaxStack()
		this.maxLocals = codeAttibute.MaxLocals()
		this.code = codeAttibute.Code()
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
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (this *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		this.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

// 从常量池种加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.slotId
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			//todo
			panic("todo")
		}
	}
}
