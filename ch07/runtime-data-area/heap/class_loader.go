package heap

import (
	"fmt"
	"go-jvm/ch07/classfile"
	"go-jvm/ch07/classpath"
)

// 类加载器
// cp字段保存ClassPath指针
// classMap保存已经加载的类数据
type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class
}

// 创建classLoader实例
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

// 把数据加载到方法区
func (this *ClassLoader) LoadClass(name string) *Class {
	//先查找类是否被加载--避免重复加载
	if class, ok := this.classMap[name]; ok {
		return class
	}
	return this.loadNonArrayClass(name)
}

// 加载数组类
func (this *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := this.readClass(name)
	class := this.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

// 类的加载分为三个步骤，加载，解析，链接
func (this *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := this.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException:" + name)
	}
	return data, entry
}

func (this *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = this
	resolveSuperClass(class)
	resolveInterfaces(class)
	this.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 计算静态函数变量,同时给予编号
func calStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 计算实例字段的个数，同时给予编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}
