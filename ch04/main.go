package main

import (
	"fmt"
	"go-jvm/ch03/classfile"
	"go-jvm/ch03/classpath"
	runtime_data_area "go-jvm/ch04/runtime-data-area"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	frame := runtime_data_area.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testOperandStack(ops *runtime_data_area.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(11601245522)
	ops.PushLong(-11601245522)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}

func testLocalVars(vs runtime_data_area.LocalVars) {
	vs.SetInt(0, 100)
	vs.SetInt(1, -100)
	vs.SetLong(2, 11601245522)
	vs.SetLong(4, -11601245522)
	vs.SetFloat(6, 3.1415926)
	vs.SetDouble(7, 2.71828182845)
	vs.SetRef(9, nil)
	println(vs.GetInt(0))
	println(vs.GetInt(1))
	println(vs.GetLong(2))
	println(vs.GetLong(4))
	println(vs.GetFloat(6))
	println(vs.GetDouble(7))
	println(vs.GetRef(9))
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classdata, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classdata)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version:%v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count:%v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags:0x%x\n", cf.AccessFlags())
	fmt.Printf("this class:%v\n", cf.ClassName())
	fmt.Printf("super class:%v\n", cf.SuperClassName())
	fmt.Printf("interfaces:%v\n", cf.InterfaceNames())
	fmt.Printf("fields count:%v", len(cf.Fields()))
	for _, m := range cf.Fields() {
		fmt.Printf("\t%s\n", m.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
