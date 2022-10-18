package heap

import "go-jvm/ch09-2/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (this *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	this.accessFlags = memberInfo.AccessFlags()
	this.name = memberInfo.Name()
	this.descriptor = memberInfo.Descriptor()
}

func (this *ClassMember) IsPublic() bool {
	return 0 != this.accessFlags&ACC_PUBLIC
}
func (this *ClassMember) IsPrivate() bool {
	return 0 != this.accessFlags&ACC_PRIVATE
}
func (this *ClassMember) IsProtected() bool {
	return 0 != this.accessFlags&ACC_PROTECTED
}
func (this *ClassMember) IsStatic() bool {
	return 0 != this.accessFlags&ACC_STATIC
}
func (this *ClassMember) IsFinal() bool {
	return 0 != this.accessFlags&ACC_FINAL
}
func (this *ClassMember) IsSynthetic() bool {
	return 0 != this.accessFlags&ACC_SYNTHETIC
}

// getters
func (this *ClassMember) Name() string {
	return this.name
}
func (this *ClassMember) Descriptor() string {
	return this.descriptor
}
func (this *ClassMember) Class() *Class {
	return this.class
}

// jvms 5.4.4
func (this *ClassMember) isAccessibleTo(d *Class) bool {
	if this.IsPublic() {
		return true
	}
	c := this.class
	if this.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	if !this.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	return d == c
}
