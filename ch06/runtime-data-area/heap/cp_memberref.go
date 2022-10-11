package heap

import "go-jvm/ch06/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (this *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	this.className = refInfo.ClassName()
	this.name, this.descriptor = refInfo.NameAndDescriptor()
}

func (this *MemberRef) Name() string {
	return this.name
}
func (this *MemberRef) Descriptor() string {
	return this.descriptor
}
