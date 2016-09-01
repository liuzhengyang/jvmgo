package classfile

import "fmt"

type ClassFile struct {
	magic uint32
	minor_version uint16
	major_version uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r. (error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic")
	}
	self.magic = magic
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minor_version = reader.readUint16()
	self.major_version = reader.readUint16()
	switch self.major_version {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minor_version == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minor_version
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.major_version
}

func (self *ClassFile) ConstantPool() ConstantPool {

}

func (self *ClassFile) AccessFlags() uint16 {

}

func (self *ClassFile) Fields() []*MemberInfo {

}

func (self *ClassFile) Methods() []*MemberInfo {

}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""   // only java.lang.Object doesn't have super class
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
