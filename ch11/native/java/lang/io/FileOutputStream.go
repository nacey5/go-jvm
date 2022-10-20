package io

import (
	"go-jvm/ch11/runtime_data_area"
	"os"
)

// private native void writeBytes(byte b[],int off,int len,boolean append)
func writeBytes(frame runtime_data_area.Frame) {
	vars := frame.LocalVars()
	//this:=vars.getRef(0)
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	//append:=vars.GetBoolean(4)
	jBytes := b.Data().([]int8)
	goBytes := castInt8sToUint8s(jBytes)
	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}
