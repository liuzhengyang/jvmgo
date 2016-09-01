package classfile

type ConstantPool []ConstantInfo

type ConstantInfo interface {
	readInifo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantPoolInfo(tag, cp)
	c.readInifo(reader)
	return c
}

func newConstantPoolInfo(tag uint8, cp ConstantPool) ConstantInfo {

}
