package heap

// 减少重复代码
// cp存放符号引用所在的运行时常量池指针
// className存放完全限定名
// class存放缓存类的指针
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (this *SymRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolveClassRef()
	}
	return this.class
}

// 如果符号已经被引用，直接返回类指针，否则调用resolveClassRef()方法解析
func (this *SymRef) resolveClassRef() {
	d := this.cp.class
	c := d.loader.LoadClass(this.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.class = c
}
