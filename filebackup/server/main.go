package main

import (
	"github.com/Wayne-Liu/hello-grpc/filebackup"
	"os"
)

func main()  {
	base := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\gen\\hello.pb.go"
	//fileList := &pb.FileList{}
	//
	//filebackup.GetFileList(base, "",fileList)
	//
	//println(len(fileList.File))

	file, err  := os.Open(base)
	if err != nil {
		println(err)
	}
	str := filebackup.GetSha256hash(file)
	println(str)

	base2 := "E:\\backup\\model\\gen\\hello.pb.go"
	file2, err  := os.Open(base2)
	if err != nil {
		println(err)
	}
	str2 := filebackup.GetSha256hash(file2)
	println(str2)
}
