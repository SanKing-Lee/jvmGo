package classfile

// SourceFile属性
type SourceFileAttribute struct {
	cp              ConstantPool // 当前的常量
	sourceFileIndex uint16       // 源文件索引
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
