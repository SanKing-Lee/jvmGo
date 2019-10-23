package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/*
 *@description: 找出目录中的所有jar文件并生成新的zipEntry放入到一个compositeEntry中去
 *@params: path: 指定路径
 *@return: 该路径下所有的jar文件组成的CompositeEntry
 *@version:
  0.0.1 初始化
*/
func newWildcardEntry(path string) CompositeEntry {
	// 去掉路径末尾的'*'
	baseDir := path[:len(path)-1]
	// 新建组合入口
	compositeEntry := []Entry{}
	// 闭包函数用于文件树的步进，在当前的基目录下找到所有的jar文件，并加入到compositeEntry中
	walkFn := func(path string, info os.FileInfo, err error) error {
		// 在步进的过程中出现了错误，则返回该错误并停止步进
		if err != nil {
			return err
		}
		// 如果当前的文件是一个目录而且当前的路径不是基目录，则跳过该目录
		// 即只搜索当前目录下的文件，不搜索当前目录下的目录中的文件
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 如果当前的路径是以jar结尾，说明这是一个jar文件，新生成一个zipEntry来表示它
		// 并将这个zipEntry放入到compositeEntry中去
		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 遍历baseDir创建ZipEntry
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
