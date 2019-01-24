package filebackup

import (
	"os"
	"github.com/lunny/log"
	"io/ioutil"
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"crypto/md5"
	"fmt"
	"io"
	"encoding/hex"
)

func GetFileList(path string, fileList *pb.FileList) {

	rd, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			GetFileList(path + fi.Name() + "/",fileList)
		} else {
			pbfile := &pb.File{}
			pbfile.Name = fi.Name()
			pbfile.Size = fi.Size()
			file,err := os.Open(path + fi.Name())
			if err != nil {
				log.Fatal("read file error:",err)
			}
			pbfile.Hash = GetMD5hash(file)
			pbfile.Path = path + fi.Name()

			fileList.File = append(fileList.File,pbfile)
		}
	}
}

func GetMD5hash(file *os.File) string {
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, file); err != nil {
		fmt.Println("Copy", err)

	}

	md := md5hash.Sum(nil)
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

func CheckSliceAaddOrUpdate(listA pb.FileList,listB pb.FileList) *pb.FileList {
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

func BackupFileBinary(list pb.FileList) *pb.FileBinary {
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
		ioutil.WriteFile(fi.Path,buf,0644)
		fileValue.Binary = buf
		fileBinary.Files = append(fileBinary.Files,fileValue)
	}
	return fileBinary
}