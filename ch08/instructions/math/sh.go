package math

// 位移指令
import (
	"go-jvm/ch08/instructions/base"
	runtime_data_area "go-jvm/ch08/runtime-data-area"
)

// ISHL int算术左位移
type ISHL struct {
	base.NoOperandsInstruction
}

// ISHR int算术右位移
type ISHR struct {
	base.NoOperandsInstruction
}

// IUSHR int 逻辑右位移
type IUSHR struct {
	base.NoOperandsInstruction
}

// LSHL long 算术左位移
type LSHL struct {
	base.NoOperandsInstruction
}

// LSHR long算术右位移
type LSHR struct {
	base.NoOperandsInstruction
}

// LUSHR long逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}

// Execute 算术左移
func (this *ISHL) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	//指出要位移多少位
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//int变量只有32位，只取出v2的前五个bit就足够表示位移数
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// Execute 算术右移
func (this *ISHR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	//指出要位移多少位
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//int变量只有32位，只取出v2的前五个bit就足够表示位移数
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// Execute 算术右移
func (this *LSHR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	//指出要位移多少位
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//long变量只有64位，只取出v2的前6个bit就足够表示位移数
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Execute 算术左移
func (this *LSHL) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	//指出要位移多少位
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//long变量只有64位，只取出v2的前6个bit就足够表示位移数
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// Execute 逻辑右移
func (this *IUSHR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	//指出要位移多少位
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// Execute 逻辑右移
func (this *LUSHR) Execute(frame *runtime_data_area.Frame) {
	stack := frame.OperandStack()
	//指出要位移多少位
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
