package heap

import "go-jvm/ch07/classfile"

// 方法符号引用
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

func (this *MethodRef) ResolvedMethod() *Method {
	if this.method == nil {
		this.resolveMethodRef()
	}
	return this.method
}

// jvms8 5.4.3.3
func (this *MethodRef) resolveMethodRef() {
	//class := self.Class()
	// todo
}
