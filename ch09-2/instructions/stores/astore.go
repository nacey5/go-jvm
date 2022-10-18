package stores

import "go-jvm/ch09-2/instructions/base"
import "go-jvm/ch09-2/rtda"

// Store reference into local variable
type ASTORE struct{ base.Index8Instruction }

func (this *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, uint(this.Index))
}

type ASTORE_0 struct{ base.NoOperandsInstruction }

func (this *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperandsInstruction }

func (this *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperandsInstruction }

func (this *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperandsInstruction }

func (this *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
