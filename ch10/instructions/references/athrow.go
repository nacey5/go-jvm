package references

import (
	"go-jvm/ch10/instructions/base"
	"go-jvm/ch10/runtime_data_area"
	"go-jvm/ch10/runtime_data_area/heap"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (this *ATHROW) Execute(frame *runtime_data_area.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}
	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handlerUncaughtException(thread, ex)
	}
}

func handlerUncaughtException(thread *runtime_data_area.Thread, ex *heap.Object) {
	thread.ClearStack()
	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ":" + goMsg)
	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat" + ste.String())
	}
}

func findAndGotoExceptionHandler(thread *runtime_data_area.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1
		handlerPc := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPc > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPc)
			return true
		}
		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}
