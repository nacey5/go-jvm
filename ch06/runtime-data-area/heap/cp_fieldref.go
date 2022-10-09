package heap

import "go-jvm/ch06/classfile"

// 字段符号引用
type FieldRef struct {
	MemberRef
	field *Field
}

// field字段缓存解析后的字段指针，newFieldRef()方法创建Field实例
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}
