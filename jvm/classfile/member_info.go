package classfile

/*
 *@description: 用于描述方法和字段的结构体
 *@version:
  0.0.1 初始化
*/
type MemberInfo struct {
	cp              ConstantPool    // 常量池
	accessFlags     uint16          // 访问标志
	nameIndex       uint16          // 对常量表的一个字符串有效索引，要么表示一个特殊方法名，要么表示方法的有效非限定名
	descriptorIndex uint16          // 对常量表的有效索引，表示一个有效的方法的描述符
	attributes      []AttributeInfo // 某个方法的任意个关联索引
}

/*
 *@description: 从Class文件中读取方法或字段的相应信息
 *@params: reader: 阅读器，cp: 当前的常量池
 *@return: 方法表或字段表
 *@version:
  0.0.1 初始化
*/
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 读取数量
	memberCount := reader.readUint16()
	// 初始化表
	members := make([]*MemberInfo, memberCount)
	// 遍历读取填充
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

/*
 *@description: 从接下来的字节流中逐个读取每一个方法或字段
 *@params: reader: 阅读器， cp：当前的常量池
 *@return: *MemberInfo: 从Class文件中读取到的一个新的方法描述或字段描述
 *@version:
  0.0.1 初始化
*/
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

//*************************************************************getter*********************************************************//
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
