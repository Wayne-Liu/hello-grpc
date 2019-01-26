package filebackup

import (
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"testing"
)

func Test_BackupFileBinary(t *testing.T)  {
	fileList := &pb.FileList{}
	path := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\gen\\"
	GetFileList(path, "",fileList)

	fileBinary := BackupFileBinary(fileList,path)

	if len(fileBinary.Files) >1 {
		t.Log("读取列表数为:",len(fileBinary.Files))
	} else {
		t.Error("读取失败不通过")
	}
}

func Test_WriteFile(t *testing.T)  {
	fileList := &pb.FileList{}
	path := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\gen\\"
	GetFileList(path, "",fileList)

	fileBinary := BackupFileBinary(fileList,path)
	t.Log("读取列表数为:",len(fileBinary.Files))
	basePath := "E:\\backup\\model\\"

	WriteFile(basePath, fileBinary)
	t.Log("写入文件条数：")

}
