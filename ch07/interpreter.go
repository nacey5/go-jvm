package main

import (
	"fmt"
	"go-jvm/ch07/instructions"
	"go-jvm/ch07/instructions/base"
	runtime_data_area "go-jvm/ch07/runtime-data-area"
	"go-jvm/ch07/runtime-data-area/heap"
)

// 解释器
// interpret 的参数是MemberInfo指针
func interpret(method *heap.Method, logInst bool) {
	//创建一个Thread实例，然后创建一个帧并把它推入java虚拟机栈中
	thread := runtime_data_area.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(thread)
	loop(thread, logInst)
}

func loop(thread *runtime_data_area.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPc(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.Pc())
		if logInst {
			logInstruction(frame, inst)
		}
		//execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *runtime_data_area.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().Pc()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

// 最后的指令都是return，而目前并没有实现return，所以必然会发生错误，所以需要recover
func catchErr(thread *runtime_data_area.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *runtime_data_area.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">>pc:%4d %v.%v%v\n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
