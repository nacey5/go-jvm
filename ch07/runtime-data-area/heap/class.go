package heap

import (
	"go-jvm/ch07/classfile"
	"strings"
)

// 表示要放进方法区的类
type Class struct {
	//类的访问标志
	accessFlags uint16
	//this class Name
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (this *Class) IsPublic() bool {
	return 0 != this.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (this *Class) ConstantPool() *ConstantPool {
	return this.constantPool
}
func (this *Class) StaticVars() Slots {
	return this.staticVars
}

// 如果D想要访问C，必须满足：C是public或者C和D在同一个包中
func (this *Class) isAccessibleTo(other *Class) bool {
	return this.IsPublic() || this.getPackageName() == other.getPackageName()
}

func (this *Class) getPackageName() string {
	if i := strings.LastIndex(this.name, "/"); i >= 0 {
		return this.name[:i]
	}
	return ""
}

func (this *Class) NewObject() *Object {
	return newObject(this)
}

func (this *Class) GetMainMethod() *Method {
	return this.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (this *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range this.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}
