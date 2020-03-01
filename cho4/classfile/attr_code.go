package classfile

/*
Code_attrbute {
	u2	attribute_name_index;
	u4	attribute_length;
	u2	max_stack;
	u2	max_locals;
	u4	code_length;
	u1	code[code_length];
	u2	exception_table_length;
	{
		u2	start_pc;
		u2	end_pc;
		u2	handler_pc;
		u2	catck_type;
	} exception_table[exception_table_length]
	u2	attributes_count;
	attribute_info	attributes[attributes_count]
}
*/

type CodeAttribute struct {
	cp 				ConstantPool
	maxStack 		uint16
	maxLocals 		uint16
	code 			[]byte
	exceptionTable	[]*ExceptionTablEntry
	attributes 		[]AttributeInfo
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUnit16()
	self.maxLocals = reader.readUnit16()
	codeLength := reader.readUnit32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTabl() []*ExceptionTablEntry {
	return self.exceptionTable
}


type ExceptionTablEntry struct {
	startPc 	uint16
	endPc	 	uint16
	handlerPc 	uint16
	catchType 	uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTablEntry {
	exceptionTableLength := reader.readUnit16()
	exceptionTable := make([]*ExceptionTablEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTablEntry {
			startPc:	reader.readUnit16(),
			endPc:		reader.readUnit16(),
			handlerPc:	reader.readUnit16(),
			catchType:	reader.readUnit16(),
		}
	}
	return exceptionTable
}


func (self *ExceptionTablEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTablEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTablEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTablEntry) CatchType() uint16 {
	return self.catchType
}


