package helper

import (
	"io/ioutil"
	"os"
	"strings"
)

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPath string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if strings.HasPrefix(fi.Name(),".")==false {
			if fi.IsDir() {
				dirs = append(dirs, dirPath+PthSep+fi.Name())
				newfiles, _ := GetAllFiles(dirPath + PthSep + fi.Name())
				files = append(files, newfiles...)
			} else {
				if  strings.HasSuffix(fi.Name(),".go")==true {
					files = append(files, dirPath+PthSep+fi.Name())
				}
			}
		}

	}

	return files, nil
}

