package heap

import "go-jvm/ch07/classfile"

// 接口方法符号引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (this *InterfaceMethodRef) ResolveInterfaceMethod() *Method {
	if this.method == nil {
		this.resolveInterfaceMethodRef()
	}
	return this.method
}
func (this *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := this.cp.class
	c := this.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, this.name, this.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.method = method
}
