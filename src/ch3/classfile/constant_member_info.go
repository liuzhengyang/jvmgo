package classfile

type ConstantMemberInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberInfo) NameAndDescriptor() string {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldInfo struct {
	ConstantMemberInfo
}

type ConstantMethodInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceInfo struct {
	ConstantMemberInfo
}
