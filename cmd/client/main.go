package main

import (
	"fmt"
	"log"


	pb "github.com/KadirbekSharau/file-to-ipld/internal/ipld-grpc/api/ipld"
	"github.com/KadirbekSharau/file-to-ipld/internal/ipld-grpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	FileStorageDiskPath = "disk/files/"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewIPLDServiceClient(conn)
	client := client.NewIPLDServiceClient(grpcClient, FileStorageDiskPath)


	fmt.Print("Enter the path of the file you want to upload: ")

	cid, err := client.UploadFile("/Users/kadirbeksharau/Desktop/aaa.txt")
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	} else {
		fmt.Printf("File uploaded successfully. CID: %s\n", cid)
	}

	fileData := client.GetFileByCID("QmTAsG4U3jwvxFFMokS8qQtgpvE6Phe6rjqG7PwinLHcqA")
	fmt.Println(string(fileData))
}