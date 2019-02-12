package filebackup

import (
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"testing"
)

const rootPath  = "/Users/liuweipeng/Desktop/document/lucene"

func Test_GetFileList(t *testing.T)  {
	fileList := &pb.FileList{}
	GetFileList(rootPath,fileList)
	if len(fileList.File) >0 {
		t.Log("fileList size is ",len(fileList.File))
	} else {
		t.Error("read file err")
	}
}

func Test_CheckSliceAaddOrUpdate(t *testing.T)  {
	fileListA := &pb.FileList{}
	GetFileList(rootPath,fileListA)
	fileListB := &pb.FileList{}
	targetPath := "E:\\backup\\cicd\\"
	GetFileList(targetPath,fileListB)
	fileListUp := CheckSliceAaddOrUpdate(fileListA,fileListB)
	t.Log(len(fileListUp.File))

}

func Test_BackupFileBinary(t *testing.T)  {
	fileList := &pb.FileList{}
	//path := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\gen\\"
	GetFileList(rootPath,fileList)

	fileBinary := BackupFileBinary(fileList,rootPath)

	if len(fileBinary.Files) >1 {
		t.Log("读取列表数为:",len(fileBinary.Files))
	} else {
		t.Error("读取失败不通过")
	}
}

func Test_WriteFile(t *testing.T)  {
	fileList := &pb.FileList{}
	//path := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\gen\\"
	GetFileList(rootPath, fileList)

	fileBinary := BackupFileBinary(fileList,rootPath)
	t.Log("读取列表数为:",len(fileBinary.Files))
	basePath := "E:\\backup\\cicd\\"

	WriteFile(basePath, fileBinary)
	t.Log("写入文件条数：")

}
