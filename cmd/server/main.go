package main

import (
	"log"
	"net"

	pb "github.com/KadirbekSharau/file-to-ipld/internal/ipld-grpc/api/ipld"
	"github.com/KadirbekSharau/file-to-ipld/internal/ipld-grpc/server"
	fileipld "github.com/KadirbekSharau/file-to-ipld/pkg/file/file-ipld"
	"google.golang.org/grpc"
)

const (
	IPLDdiskPath = "disk/ipld/"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv := server.NewIPLDServiceServer(
		fileipld.NewIPLDFileManager(IPLDdiskPath),
	)
	pb.RegisterIPLDServiceServer(grpcServer, srv)

	log.Println("Starting server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}