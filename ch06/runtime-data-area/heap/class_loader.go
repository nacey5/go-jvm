package heap

import (
	"fmt"
	"go-jvm/ch06/classpath"
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
