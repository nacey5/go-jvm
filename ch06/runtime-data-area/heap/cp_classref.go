package heap

import "go-jvm/ch06/classfile"

type ClassRef struct {
	SymRef
}

// 根据class文件存储的类的常量创建实例
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
