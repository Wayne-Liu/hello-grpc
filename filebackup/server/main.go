package main

import (
	"github.com/Wayne-Liu/hello-grpc/filebackup"
	"os"
	"path/filepath"
	"fmt"
	"strings"
)

const(
	rootDir = "D:/apache-tomcat-8.0.35/webapps/docs/"
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

	base2 := "E:\\backup\\model\\hello.pb.go"
	file2, err  := os.Open(base2)
	if err != nil {
		println(err)
	}
	str2 := filebackup.GetSha256hash(file2)
	println(str2)

	//println(string(os.PathSeparator))
	
	//filepath.Walk("D:/apache-tomcat-8.0.35/webapps/docs",findFileDir)

	println(filepath.Base(rootDir))
	println(filepath.Dir(rootDir))
}



func findFileDir(path string, info os.FileInfo, err error) error {
	//ok, err := filepath.Match(`*.txt`, info.Name())
	//if ok {
	//	fmt.Println(filepath.Dir(path), info.Name())
	//	// 遇到 txt 文件则继续处理所在目录的下一个目录
	//	// 注意会跳过子目录
	//	return filepath.SkipDir
	//}
	//return err
	rootDirWin := filepath.Dir(rootDir)
	//fmt.Println("rootDirWin:"+rootDirWin)

	if !info.IsDir() {
		//fmt.Println(filepath.Dir(path), info.Name())
		relativePath := strings.Replace(path,rootDirWin,"",-1)
		fmt.Println(relativePath)
	}
	return nil
}