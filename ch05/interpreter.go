package main

import (
	"fmt"
	"go-jvm/ch05/classfile"
	"go-jvm/ch05/instructions"
	"go-jvm/ch05/instructions/base"
	runtime_data_area "go-jvm/ch05/runtime-data-area"
)

// 解释器
// interpret 的参数是MemberInfo指针
func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	//创建一个Thread实例，然后创建一个帧并把它推入java虚拟机栈中
	thread := runtime_data_area.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

func loop(thread *runtime_data_area.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPc(pc)

		//decode
		reader.Reset(bytecode, pc)
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.Pc())

		//execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

// 最后的指令都是return，而目前并没有实现return，所以必然会发生错误，所以需要recover
func catchErr(frame *runtime_data_area.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
