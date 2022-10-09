package classfile

// LineNumberTableAttribute 行号信息
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

// LineNumberTableEntry 行号实体信息
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

// 读取行信息
func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberLength := reader.readUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, lineNumberLength)
	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

// 得到行数
func (this *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(this.lineNumberTable) - 1; i >= 0; i-- {
		entry := this.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
