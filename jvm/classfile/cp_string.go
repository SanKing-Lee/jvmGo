package classfile

type ConstantStringInfo struct {
	cp          ConstantPool // 当前的常量池
	stringIndex uint16       // 当前string指向的utf8在常量池中的索引
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
