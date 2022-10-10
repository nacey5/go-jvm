package heap

// class 存放对象的class指针，一个存放对象实例
type Object struct {
	class  *Class
	fields Slots
}
