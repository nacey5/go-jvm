package atomic

import "go-jvm/ch11-2/native"
import "go-jvm/ch11-2/runtime_data_area"

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *runtime_data_area.Frame) {
	frame.OperandStack().PushBoolean(false)
}
