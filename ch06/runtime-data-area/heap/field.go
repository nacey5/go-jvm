package heap

import "go-jvm/ch06/classfile"

type Field struct {
	ClassMember
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
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
