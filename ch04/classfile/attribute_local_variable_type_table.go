package classfile

// LocalVariableTypeTableAttribute 局部表变量类型
type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (this *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader) {
	localVariableTypeTableLength := reader.readUint16()
	this.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range this.localVariableTypeTable {
		this.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{
			startPc:        reader.readUint16(),
			length:         reader.readUint16(),
			nameIndex:      reader.readUint16(),
			signatureIndex: reader.readUint16(),
			index:          reader.readUint16(),
		}
	}
}
