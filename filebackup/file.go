package filebackup

import (
	"os"
	"github.com/lunny/log"
	"io/ioutil"
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"fmt"
	"io"
	"encoding/hex"
	"strings"
	"crypto/sha256"
	"path/filepath"
)

func GetFileList(path string, fileList *pb.FileList) error{
	//currentDir := filepath.Dir(path)
	//dir := filepath.Base(path)
	//parentDir := strings.Replace(currentDir,dir,"",-1)
	//判断目录是否存在，如果不存在在返回空
	bool, err := PathExists(path)
	if !bool {
		return err
	}

	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {

		if !info.IsDir(){
			pbfile := &pb.File{}
			pbfile.Name = info.Name()
			pbfile.Size = info.Size()
			file,err := os.Open(filePath)
			if err != nil {
				log.Fatal("read file error:",err)
			}
			pbfile.Hash = GetSha256hash(file)
			relativePath := strings.Replace(filePath,path,"",-1)
			pbfile.Path = relativePath

			fileList.File = append(fileList.File,pbfile)
		}
		return err
	})
	return err
}


func GetSha256hash(file *os.File) string {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("Copy", err)

	}

	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CheckSliceAaddOrUpdate(listA *pb.FileList,listB *pb.FileList) *pb.FileList {
	addFileList := &pb.FileList{}
	fileA := listA.File
	fileB := listB.File
	for _, valueA := range fileA {
		hasBackup := false
		for i := 0; i < len(fileB); i++ {
			if valueA.Hash == fileB[i].Hash {
				hasBackup = true
			}
		}
		if !hasBackup {
			//valueA.Type = "update"
			addFileList.File = append(addFileList.File,valueA)
		}
	}
	return addFileList
}

func BackupFileBinary(list *pb.FileList,basePath string) *pb.FileBinary {
	fileBinary := &pb.FileBinary{}
	files := list.File
	for _, fi := range files {
		fileValue := &pb.FileValue{}
		fileValue.Path = fi.Path
		//file,err := os.Open(fi.Path)
		//if err != nil {
		//	log.Fatal("read file err :",err)
		//}
		//buf := make([]byte,fi.Size)
		buf,err := ioutil.ReadFile(basePath+fi.Path)
		if err != nil {
			log.Fatal("read file error :",err)
		}
		fileValue.Binary = buf
		fileBinary.Files = append(fileBinary.Files,fileValue)

	}
	return fileBinary
}

func WriteFile(basePath string, value *pb.FileBinary)  {
	fileValues := value.Files
	log.Println("file numbers is :", len(fileValues))
	for _, fi := range fileValues{
		filePath := basePath + fi.Path
		err := os.MkdirAll(getDir(filePath), 0644)
		if err != nil {
			log.Fatal(err)
		}

		w, err := os.Create(filePath)
		if err != nil {
			log.Fatal("Create file error:",err)
		}

		w.Close()

		err = ioutil.WriteFile(filePath,fi.Binary,0644)
		if err != nil {
			log.Fatal("Write file error:",err)
		}
	}
}

func getDir(path string) string {
	if strings.Contains(path,"/") {
		return subString(path, 0, strings.LastIndex(path, "/"))
	} else {
		return subString(path, 0, strings.LastIndex(path, "\\"))
	}

}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}