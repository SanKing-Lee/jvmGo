package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableEntry
}

type LocalVariableEntry struct {
	startPc         uint16 // 局部变量开始的指令
	length          uint16 // 局部变量长度范围
	nameIndex       uint16 // 局部变量的名称在常量池中的索引
	descriptorIndex uint16 // 局部变量的描述符在常量池中的索引
	index           uint16 // 局部变量在栈帧中的索引
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
