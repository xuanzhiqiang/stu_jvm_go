package classfile

type ConstantMemberrefInfo struct {
	cp 					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUnit16()
	self.nameAndTypeIndex = reader.readUnit16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getUtf8(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}


type ConstantFieldrefInfo struct {ConstantMemberrefInfo}
type ConstantMethodrefInfo struct {ConstantMemberrefInfo}
type ConstantInterfaceMethodrefInfo struct {ConstantMemberrefInfo}