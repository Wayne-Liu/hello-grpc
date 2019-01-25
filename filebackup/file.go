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
)

func GetFileList(path string,relative string, fileList *pb.FileList) {

	rd, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			relative += fi.Name() + "/"
			GetFileList(path + fi.Name() + "/", relative,fileList)
		} else {
			pbfile := &pb.File{}
			pbfile.Name = fi.Name()
			pbfile.Size = fi.Size()
			file,err := os.Open(path + fi.Name())
			if err != nil {
				log.Fatal("read file error:",err)
			}
			pbfile.Hash = GetSha256hash(file)
			pbfile.Path = relative + fi.Name()

			fileList.File = append(fileList.File,pbfile)
		}
	}
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
		buf := make([]byte,fi.Size)
		ioutil.ReadFile(basePath+fi.Path)
		fileValue.Binary = buf
		fileBinary.Files = append(fileBinary.Files,fileValue)
	}
	return fileBinary
}

func WriteFile(basePath string, value *pb.FileBinary)  {
	fileValues := value.Files
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
	return subString(path, 0, strings.LastIndex(path, "/"))
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