package extended

import (
	"go-jvm/ch06/instructions/base"
	"go-jvm/ch06/instructions/loads"
	"go-jvm/ch06/instructions/math"
	"go-jvm/ch06/instructions/stores"
	runtime_data_area "go-jvm/ch06/runtime-data-area"
)

// 扩展类指令
type WIDE struct {
	modifiedInstruction base.Instruction
}

// 加载指令和存储指令都只有一个操作数，需要扩展成2字节
func (this *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	//iload
	case 0x15:
		{
			inst := &loads.ILOAD{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//lload
	case 0x16:
		{
			inst := &loads.LLOAD{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//fload
	case 0x17:
		{
			inst := &loads.FLOAD{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//dload
	case 0x18:
		{
			inst := &loads.DLOAD{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//aload
	case 0x19:
		{
			inst := &loads.ALOAD{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//istore
	case 0x36:
		{
			inst := &stores.ISTORE{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//lstore
	case 0x37:
		{
			inst := &stores.LSTORE{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//fstore
	case 0x38:
		{
			inst := &stores.FSTORE{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//dstore
	case 0x39:
		{
			inst := &stores.DSTORE{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//astore
	case 0x3a:
		{
			inst := &stores.ASTORE{}
			inst.Index = uint(reader.ReadUint16())
			this.modifiedInstruction = inst
		}
		//iinc
	case 0x84:
		{
			inst := &math.IINC{}
			inst.Index = uint(reader.ReadUint16())
			inst.Const = int32(reader.ReadInt16())
			this.modifiedInstruction = inst
		}
		//ret
	case 0xa9:
		{
			panic("Unsupported opcode:0xa9")
		}
	}
}

func (this *WIDE) Execute(frame *runtime_data_area.Frame) {
	this.modifiedInstruction.Execute(frame)
}
