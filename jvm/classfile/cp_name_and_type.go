package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16 // 名称索引
	descriptorIndex uint16 // 描述符索引
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
