package main

import (
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"github.com/Wayne-Liu/hello-grpc/filebackup"
)

func main()  {
	path := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\"
	fileList := &pb.FileList{}

	filebackup.GetFileList(path, fileList)

	println(len(fileList.File))


}
