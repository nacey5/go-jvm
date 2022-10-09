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
