package classfile

type CodeAttribute struct {
	cp             ConstantPool           // 当前的常量池
	maxStack       uint16                 // 栈的最大深度
	maxLocals      uint16                 // 本地变量的最大数量
	code           []byte                 // 表示代码段的字节码
	exceptionTable []*ExceptionTableEntry // 异常表
	attributes     []AttributeInfo        // 属性表
}

type ExceptionTableEntry struct {
	startPc   uint16 // 异常处理的起始pc地址
	endPc     uint16 // 异常处理的终止pc地址
	handlerPc uint16 // 异常处理的处理PC地址，即catch段
	catchType uint16 // 异常处理需要捕获的异常类型
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
