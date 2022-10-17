package heap

import "go-jvm/ch08-2/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (this *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		this.maxStack = codeAttr.MaxStack()
		this.maxLocals = codeAttr.MaxLocals()
		this.code = codeAttr.Code()
	}
}

func (this *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(this.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		this.argSlotCount++
		if paramType == "J" || paramType == "D" {
			this.argSlotCount++
		}
	}
	if !this.IsStatic() {
		this.argSlotCount++ // `this` reference
	}
}

func (this *Method) IsSynchronized() bool {
	return 0 != this.accessFlags&ACC_SYNCHRONIZED
}
func (this *Method) IsBridge() bool {
	return 0 != this.accessFlags&ACC_BRIDGE
}
func (this *Method) IsVarargs() bool {
	return 0 != this.accessFlags&ACC_VARARGS
}
func (this *Method) IsNative() bool {
	return 0 != this.accessFlags&ACC_NATIVE
}
func (this *Method) IsAbstract() bool {
	return 0 != this.accessFlags&ACC_ABSTRACT
}
func (this *Method) IsStrict() bool {
	return 0 != this.accessFlags&ACC_STRICT
}

// getters
func (this *Method) MaxStack() uint {
	return this.maxStack
}
func (this *Method) MaxLocals() uint {
	return this.maxLocals
}
func (this *Method) Code() []byte {
	return this.code
}
func (this *Method) ArgSlotCount() uint {
	return this.argSlotCount
}
