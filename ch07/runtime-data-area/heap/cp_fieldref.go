package heap

import "go-jvm/ch07/classfile"

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

func (this *FieldRef) ResolvedField() *Field {
	if this.field == nil {
		this.ResolveFieldRef()
	}
	return this.field
}

// 根据虚拟机贵方将其翻译成ResolveFieldRef()
func (this *FieldRef) ResolveFieldRef() {
	d := this.cp.class
	c := this.ResolvedClass()
	field := lookupField(c, this.name, this.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.field = field
}

// 如果D类想要访问C类某个字段，就要先解析符号引用得到类C，再根据字段名进行查找
func lookupField(c *Class, name string, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}
