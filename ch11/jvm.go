package main

import (
	"fmt"
	"go-jvm/ch11/classpath"
	"go-jvm/ch11/instructions/base"
	"go-jvm/ch11/runtime_data_area"
	"go-jvm/ch11/runtime_data_area/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *runtime_data_area.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  runtime_data_area.NewThread(),
	}
}

func (this *JVM) start() {
	this.initVM()
	this.execMain()
}

func (this *JVM) initVM() {
	vmClass := this.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(this.mainThread, vmClass)
	interpret(this.mainThread, this.cmd.verboseInstFlag)
}

func (this *JVM) execMain() {
	className := strings.Replace(this.cmd.class, ".", "/", -1)
	mainClass := this.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", this.cmd.class)
		return
	}
	argsArr := this.createArgsArray()
	frame := this.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr) //给main方法传递args参数
	this.mainThread.PushFrame(frame)
	interpret(this.mainThread, this.cmd.verboseInstFlag)
}

func (this *JVM) createArgsArray() *heap.Object {
	stringClass := this.classLoader.LoadClass("java/lang/String")
	argsLen := uint(len(this.cmd.args))
	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range this.cmd.args {
		jArgs[i] = heap.JString(this.classLoader, arg)
	}
	return argsArr
}
