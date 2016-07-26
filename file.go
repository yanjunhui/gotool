package gotool

import (
	"os"
	"bufio"
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


//遍历目录获取文件列表
type Dir struct{
	Name string
	Files []os.FileInfo
}
var dirInfo []Dir

func RangeDir(dir string)([]Dir, error){
	var files []os.FileInfo
	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return dirInfo, err
	}

	for _, f := range filesInfo {
		if f.IsDir(){
			_, err = RangeDir(dir + "/" + f.Name())

		}else {
			files = append(files, f)
		}

	}
	newDir := Dir{
		Name:dir,
		Files:files,
	}
	dirInfo = append(dirInfo, newDir)
	return  dirInfo, err
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