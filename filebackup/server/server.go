package main

import (
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"context"
	pb "github.com/Wayne-Liu/hello-grpc/filebackup/file"
	"log"
	"github.com/Wayne-Liu/hello-grpc/filebackup"
	"os"
	"path/filepath"
)

const (
	//port = ":80"
	//basePath = "/root/backupfile/"
	port = ":50015"
	basePath = "E:\\backup\\"
)

// server is used to implement helloworld.GreeterServer.
type server struct{
	fileBasePath string
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetRemoteFileList(ctx context.Context, in *pb.HelloRequest) (*pb.FileList, error) {
	log.Printf("FilePath is: %v", in.Path)
	fileBasePath := filepath.Dir(basePath) + string(os.PathSeparator) + in.Path+ string(os.PathSeparator)
	s.fileBasePath = fileBasePath

	if !Exists(fileBasePath){
		os.Mkdir(fileBasePath,0644)
	}
	fileList := &pb.FileList{}
	filebackup.GetFileList(fileBasePath,fileList)

	return fileList, nil
}

func (s *server) SendFiles(ctx context.Context, in *pb.FileBinary) (*pb.HelloReply, error) {
	log.Println("fileBasePath is :"+s.fileBasePath)
	filebackup.WriteFile(s.fileBasePath,in)
	msg := "add file numbers: " + string(len(in.Files))
	return &pb.HelloReply{Message: msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("server is starting...")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	println("server is started!")
}

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
