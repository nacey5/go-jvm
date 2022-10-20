package misc

import (
	"go-jvm/ch11/instructions/base"
	"go-jvm/ch11/native"
)
import "go-jvm/ch11/runtime_data_area"

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

// private static native void initialize();
// ()V
func initialize(frame *runtime_data_area.Frame) { //取消掉上版本的hack hack: just make VM.savedProps nonempty
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
