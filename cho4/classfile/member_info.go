package classfile

/**
field_info {
	u2				access_flags;
	u2				name_index;
	u2				descriptor_index;
	u2				attributes_count;
	attribute_info	attributes[attributes_count];
}
method_info {
	u2				access_flags;
	u2				name_index;
	u2				descriptor_index;
	u2				attributes_count;
	attribute_info	attributes[attributes_count];
}
*/


type MemberInfo struct {
	cp 					ConstantPool
	accessFlags			uint16
	nameIndex 			uint16
	descriptorIndex 	uint16
	attributes			[]AttributeInfo
}

// read field or method table
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUnit16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:					cp,
		accessFlags:		reader.readUnit16(),
		nameIndex:			reader.readUnit16(),
		descriptorIndex:	reader.readUnit16(),
		attributes:			readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}