package base

import (
	runtime_data_area "go-jvm/ch07/runtime-data-area"
	"go-jvm/ch07/runtime-data-area/heap"
)

func InitClass(thread *runtime_data_area.Thread, class *heap.Class) {
	//开始初始化--->先将初始化设置为true,避免死循环
	class.StartInit()
	//执行类的初始化
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

// 超类的执行初始化方法
func initSuperClass(thread *runtime_data_area.Thread, class *heap.Class) {
	//当不是接口的时候，获得父类
	if !class.IsInterface() {
		superClass := class.SuperClass()
		//当父类不为空且并没有被初始化过
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}

func scheduleClinit(thread *runtime_data_area.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		//exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
