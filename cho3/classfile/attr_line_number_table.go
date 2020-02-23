package classfile

type LineNumberTableAttribute struct {
	lineNumberTable 	[]*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc		uint16
	lineNumber 	uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUnit16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry {
			startPc:	reader.readUnit16(),
			lineNumber:	reader.readUnit16(),
		}
	}
}