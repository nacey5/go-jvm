package heap

import "strings"
import "go-jvm/ch09-2/classfile"

// name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
type Class struct {
	accessFlags       uint16
	name              string // thisClassName
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
	initStarted       bool
	jClass            *Object
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
func (this *Class) IsFinal() bool {
	return 0 != this.accessFlags&ACC_FINAL
}
func (this *Class) IsSuper() bool {
	return 0 != this.accessFlags&ACC_SUPER
}
func (this *Class) IsInterface() bool {
	return 0 != this.accessFlags&ACC_INTERFACE
}
func (this *Class) IsAbstract() bool {
	return 0 != this.accessFlags&ACC_ABSTRACT
}
func (this *Class) IsSynthetic() bool {
	return 0 != this.accessFlags&ACC_SYNTHETIC
}
func (this *Class) IsAnnotation() bool {
	return 0 != this.accessFlags&ACC_ANNOTATION
}
func (this *Class) IsEnum() bool {
	return 0 != this.accessFlags&ACC_ENUM
}

// getters
func (this *Class) Name() string {
	return this.name
}
func (this *Class) ConstantPool() *ConstantPool {
	return this.constantPool
}
func (this *Class) Fields() []*Field {
	return this.fields
}
func (this *Class) Methods() []*Method {
	return this.methods
}
func (this *Class) Loader() *ClassLoader {
	return this.loader
}
func (this *Class) SuperClass() *Class {
	return this.superClass
}
func (this *Class) StaticVars() Slots {
	return this.staticVars
}
func (this *Class) InitStarted() bool {
	return this.initStarted
}
func (this *Class) JClass() *Object {
	return this.jClass
}

func (this *Class) StartInit() {
	this.initStarted = true
}

// jvms 5.4.4
func (this *Class) isAccessibleTo(other *Class) bool {
	return this.IsPublic() ||
		this.GetPackageName() == other.GetPackageName()
}

func (this *Class) GetPackageName() string {
	if i := strings.LastIndex(this.name, "/"); i >= 0 {
		return this.name[:i]
	}
	return ""
}

func (this *Class) GetMainMethod() *Method {
	return this.getMethod("main", "([Ljava/lang/String;)V", true)
}
func (this *Class) GetClinitMethod() *Method {
	return this.getMethod("<clinit>", "()V", true)
}

func (this *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := this; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
}

func (this *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := this; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic &&
				field.name == name &&
				field.descriptor == descriptor {

				return field
			}
		}
	}
	return nil
}

func (this *Class) isJlObject() bool {
	return this.name == "java/lang/Object"
}
func (this *Class) isJlCloneable() bool {
	return this.name == "java/lang/Cloneable"
}
func (this *Class) isJioSerializable() bool {
	return this.name == "java/io/Serializable"
}

func (this *Class) NewObject() *Object {
	return newObject(this)
}

func (this *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(this.name)
	return this.loader.LoadClass(arrayClassName)
}

func (this *Class) JavaName() string {
	return strings.Replace(this.name, "/", ".", -1)
}

func (this *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[this.name]
	return ok
}

func (this *Class) GetInstanceMethod(name, descriptor string) *Method {
	return this.getMethod(name, descriptor, false)
}

func (this *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := this.getField(fieldName, fieldDescriptor, true)
	return this.staticVars.GetRef(field.slotId)
}
func (this *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := this.getField(fieldName, fieldDescriptor, true)
	this.staticVars.SetRef(field.slotId, ref)
}
