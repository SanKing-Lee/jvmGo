/*
 *@description:
 *@author: sean
 *@date: 2019-06-21 19:42:36
 *@version:
  0.0.1 初始化
*/
package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

/*
 *@description: 用于描述各项加载类
 *@version:
  0.0.1 初始化
*/
type Entry interface {
	readClass(className string) ([]byte, Entry, error) // 寻找和加载Class文件
	String() string                                    // 转化为字符串方法
}

/*
 *@description: 根据类路径创建一个新的类条目
 *@params: path:类路径
 *@return: 新的条目
 *@version:
  0.0.1 初始化
*/
func newEntry(path string) Entry {
	// 如果当前的类路径有分隔符，说明有多条class文件路径，生成一个组合条目
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	// 如果当前的类路径末尾有通配符
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// 如果类路径以jar或zip结尾
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasPrefix(path, ".ZIP") {
		return newZipEntry(path)
	}
	// 类路径是一个目录
	return newDirEntry(path)
}
