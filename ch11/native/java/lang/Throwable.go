package lang

import (
	"go-jvm/ch11/native"
	"go-jvm/ch11/runtime_data_area"
	"go-jvm/ch11/runtime_data_area/heap"
)

func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// public static Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *runtime_data_area.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *runtime_data_area.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func createStackTraceElement(frame *runtime_data_area.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

// 由于栈正在运行，所以必须跳过两个栈帧，具体情况要看具体实例
func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

// 记录了jvm虚拟机栈的信息
type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}
