package heap

import "go-jvm/ch10/classfile"

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            []byte
	argSlotCount    uint
	exceptionTable  ExceptionTable
	lineNumberTable *classfile.LineNumberTableAttribute
}

func (this *Method) LineNumberTable() *classfile.LineNumberTableAttribute {
	return this.lineNumberTable
}

func (this *Method) SetLineNumberTable(lineNumberTable *classfile.LineNumberTableAttribute) {
	this.lineNumberTable = lineNumberTable
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (this *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		this.maxStack = codeAttr.MaxStack()
		this.maxLocals = codeAttr.MaxLocals()
		this.code = codeAttr.Code()
		this.lineNumberTable = codeAttr.LineNumberTableAttribute()
		this.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), this.class.constantPool)
	}
}

func (this *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		this.argSlotCount++
		if paramType == "J" || paramType == "D" {
			this.argSlotCount++
		}
	}
	if !this.IsStatic() {
		this.argSlotCount++ // `this` reference
	}
}

func (this *Method) injectCodeAttribute(returnType string) {
	this.maxStack = 4 // todo
	this.maxLocals = this.argSlotCount
	switch returnType[0] {
	case 'V':
		this.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		this.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		this.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		this.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		this.code = []byte{0xfe, 0xad} // lreturn
	default:
		this.code = []byte{0xfe, 0xac} // ireturn
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

func (this *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := this.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (this *Method) GetLineNumber(pc int) int {
	if this.IsNative() {
		return -2
	}
	if this.lineNumberTable == nil {
		return -1
	}
	return this.lineNumberTable.GetLineNumber(pc)
}
