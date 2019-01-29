package filebackup

import (
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"testing"
)

const rootPath  = "/Users/liuweipeng/Desktop/document/lucene"

func Test_GetFileList(t *testing.T)  {
	fileList := &pb.FileList{}
	GetFileList(rootPath,fileList)
	if len(fileList.File) == 0 {
		t.Error("read fileList wrong!")
	} else {
		t.Log("fileList size is " , len(fileList.File))
	}
}

func Test_BackupFileBinary(t *testing.T)  {
	fileList := &pb.FileList{}
	//path := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\gen\\"
	GetFileList(rootPath, fileList)

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
	basePath := "E:\\backup\\model\\"

	WriteFile(basePath, fileBinary)
	t.Log("写入文件条数：")

}
