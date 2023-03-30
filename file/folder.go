package file

import (
	"io/fs"
	"io/ioutil"
)

type Files struct {
	List []*File
}

type File struct {
	Folder string
	Info   fs.FileInfo
}

func GetFiles(folder string, list *Files) (err error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder+"/"+file.Name(), list)
			if err != nil {
				return
			}
		} else {
			list.List = append(list.List, &File{
				Folder: folder,
				Info:   file,
			})
		}
	}
	return
}

func GetFile(folder string) (file []byte, err error) {
	return ioutil.ReadFile(folder)
}
