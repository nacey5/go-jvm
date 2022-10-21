package base

import "go-jvm/ch11-2/runtime_data_area"

func Branch(frame *runtime_data_area.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
