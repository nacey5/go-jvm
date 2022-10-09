package base

import runtime_data_area "go-jvm/ch05/runtime-data-area"

func Branch(frame *runtime_data_area.Frame, offset int) {
	pc := frame.Thread().Pc()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
