package lang

import "go-jvm/ch11-2/native"
import "go-jvm/ch11-2/runtime_data_area"

func init() {
	native.Register("java/lang/Thread", "currentThread", "()Ljava/lang/Thread;", currentThread)
	native.Register("java/lang/Thread", "setPriority0", "(I)V", setPriority0)
	native.Register("java/lang/Thread", "isAlive", "()Z", isAlive)
	native.Register("java/lang/Thread", "start0", "()V", start0)
}

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *runtime_data_area.Frame) {
	//jThread := frame.Thread().JThread()
	classLoader := frame.Method().Class().Loader()
	threadClass := classLoader.LoadClass("java/lang/Thread")
	jThread := threadClass.NewObject()

	threadGroupClass := classLoader.LoadClass("java/lang/ThreadGroup")
	jGroup := threadGroupClass.NewObject()

	jThread.SetRefVar("group", "Ljava/lang/ThreadGroup;", jGroup)
	jThread.SetIntVar("priority", "I", 1)

	frame.OperandStack().PushRef(jThread)
}

// private native void setPriority0(int newPriority);
// (I)V
func setPriority0(frame *runtime_data_area.Frame) {
	// vars := frame.LocalVars()
	// this := vars.GetThis()
	// newPriority := vars.GetInt(1))
	// todo
}

// public final native boolean isAlive();
// ()Z
func isAlive(frame *runtime_data_area.Frame) {
	//vars := frame.LocalVars()
	//this := vars.GetThis()

	stack := frame.OperandStack()
	stack.PushBoolean(false)
}

// private native void start0();
// ()V
func start0(frame *runtime_data_area.Frame) {
	// todo
}
