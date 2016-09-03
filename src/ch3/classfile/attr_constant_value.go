package classfile

type ConstantAttrValue struct {
	constantValueIndex uint16
}

func (self *ConstantAttrValue) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantAttrValue) ConstantValueIndex() uint16 {
	return self.constantValueIndex;
}