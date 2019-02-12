package main

import (
	"google.golang.org/grpc"
	"os"
	"time"
	"log"
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"context"
	"strings"
	"github.com/Wayne-Liu/hello-grpc/filebackup"
	"path/filepath"
)

const (
	//address     = "10.10.5.227:50051"
	//address = "47.52.75.214:80"
	address = "127.0.0.1:50015"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	//basePath := "C:\\Users\\liu.wp\\go\\src\\github.com\\Wayne-Liu\\hello-grpc\\model\\"
	//basePath := "D:\\BC设计资料\\"
	basePath := "D:\\apache-tomcat-8.0.35\\webapps\\docs\\"
	if len(os.Args) > 1 {
		basePath = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fileListLocal := &pb.FileList{}
	//fileListRemote := &pb.FileList{}

	filebackup.GetFileList(basePath,fileListLocal)

	helloRequest := &pb.HelloRequest{}

	helloRequest.Path = filepath.Base(basePath)
	fileListRemote, err := c.GetRemoteFileList(ctx, helloRequest)

	updateFileList := filebackup.CheckSliceAaddOrUpdate(fileListLocal,fileListRemote)

	fileBinary := filebackup.BackupFileBinary(updateFileList,basePath)

	r, err := c.SendFiles(ctx, fileBinary)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

}

func GetRootDir(basePath string) string {
	winStrs := strings.Split(basePath,"/")
	linStrs := strings.Split(basePath,"\\")
	baseDir := ""
	if len(winStrs) >2 {
		if winStrs[len(winStrs)-1] == ""{
			baseDir = winStrs[len(winStrs)-2]
		} else {
			baseDir = winStrs[len(winStrs)-1]
		}
	} else if len(linStrs) >2 {
		if linStrs[len(linStrs)-1] == "" {
			baseDir = linStrs[len(linStrs)-2]
		} else {
			baseDir = linStrs[len(linStrs)-1]
		}
	}
	return baseDir
}