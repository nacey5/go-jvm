package heap

import "go-jvm/ch07/classfile"

// 方法符号引用 ----非接口方法符号的引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodRefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

// ResolvedMethod 非接口方法符号的引用
func (this *MethodRef) ResolvedMethod() *Method {
	if this.method == nil {
		this.resolveMethodRef()
	}
	return this.method
}

// jvms8 5.4.3.3
// 如果还没有解析过符号引用，先进行解析
func (this *MethodRef) resolveMethodRef() {
	d := this.cp.class
	c := this.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangError1")
	}
	method := lookupMethod(c, this.name, this.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	//检查d类是否有权限访问这个方法
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
