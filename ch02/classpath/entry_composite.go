package classpath

import (
	"errors"
	"strings"
)

/*
 *@description: 通过组合路径生成的组合入口
 *@version:
  0.0.1 初始化
*/
type CompositeEntry []Entry

/*
 *@description: 创建一个新的组合入口
 *@params: pathList: 用文件分割符分割开的组合路径
 *@return: 组合路径中的每个路径对应的入口组合而成的组合路径
 *@version:
  0.0.1 初始化
*/
func newCompositeEntry(pathList string) CompositeEntry {
	// 初始化一个Entry切片用来放置所有之后生成的Entry
	compositeEntry := []Entry{}
	// 将多个路径组合而成的路径分割成多个路径并遍历，对每一个路径生成一个Entry并加入到Entry切片中
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 不使用指针是因为它本身就是一个切片了
/*
 *@description: 从组合路径的入口中读取指定的class文件，转换成从子路径入口寻找该class文件并读取
 *@params: className: 需要读取的指定class文件
 *@return: 读取到的数据，读取到数据的入口，错误信息
 *@version:
  0.0.1 初始化
*/
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 遍历每个子路径对应的入口
	for _, entry := range self {
		// 调用子入口的readClass方法
		data, from, err := entry.readClass(className)
		// 如果返回的err为nil说明成功读取到了，返回读取到的数据
		if err == nil {
			return data, from, nil
		}
	}
	// 遍历了一遍仍然没有找到该class文件，返回错误
	return nil, self, errors.New("class not found: " + className)
}

/*
 *@description: 获取该compositeEntry对应的字符串描述
 *@params:
 *@return: 用文件分割符分割的所有子路径对应入口的string的组合
 *@version:
  0.0.1 初始化
*/
func (self CompositeEntry) String() string {
	// 创建一个新的string切片用来存放所有子入口的string
	strs := make([]string, len(self))
	// 存放子入口的string
	for i, entry := range self {
		strs[i] = entry.String()
	}
	// 将所有子入口的string用文件分割符连接成一个新的string并返回
	return strings.Join(strs, pathListSeparator)
}
