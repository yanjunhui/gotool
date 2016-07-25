package gotool

import (
	"os"
	"bufio"
	"path/filepath"
	"io/ioutil"
)


//复制单文件(bufio缓存)
func CopyFile(srcFile string, destFile string)(size int64, err error){
	sf, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	buf := bufio.NewReader(sf)
	df, err := os.Create(destFile)
	if err != nil {
		return 0, err
	}
	return buf.WriteTo(df)
}



//删除目录
func Remove(path string) bool {
	err := os.RemoveAll(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

//判断文件是否存在
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

//遍历目录获取文件列表(一)
func RangeDirWalk(dir string)(files []string){
	filepath.Walk(dir,
		func(file string, f os.FileInfo, err error) error {
			if (f == nil) {
				return err
			}
			if f.IsDir() {
				return nil
			}
			files = append(files, file)
			return nil
		})
	return files
}


//递归复制目录
func CopyDir(srcDir string, destDir string)(err error){
	if IsExist(destDir) == false {
		os.Mkdir(destDir, 0755)
	}
	d, err := ioutil.ReadDir(srcDir)
	for _, file := range d{
		if file.IsDir(){
			path := srcDir + "/" + file.Name()
			CopyDir(path, destDir + "/" + file.Name())
		}else {
			_, err = CopyFile(srcDir + "/" + file.Name(), destDir + "/" + file.Name())
		}
	}
	return err
}