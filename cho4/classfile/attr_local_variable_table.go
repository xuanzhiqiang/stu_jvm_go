package classfile

/*
LocalVariableTable_attribute {
	u2	attribute_name_index;
	u4	attribute_length;
	u2	local_variable_table_length;
	{
		u2	startPc;
		u2	length;
		u2	name_index;
		u2	descriptor_index;
		u2	index;
	} local_variable_table[local_variable_table_length]
}
*/

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc			uint16
	length			uint16
	nameIndex		uint16
	descriptorIndex	uint16
	index			uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUnit16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry {
			startPc:			reader.readUnit16(),
			length:				reader.readUnit16(),
			nameIndex:			reader.readUnit16(),
			descriptorIndex:	reader.readUnit16(),
			index:				reader.readUnit16(),
		}
	}
}