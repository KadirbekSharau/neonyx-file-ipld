package client

import (
	"context"
	"log"
	"os"

	pb "github.com/KadirbekSharau/file-to-ipld/internal/ipld-grpc/api/ipld"
)

type IPLDServiceClient struct {
	client pb.IPLDServiceClient
	FileStoragePath string
}

func NewIPLDServiceClient(client pb.IPLDServiceClient, FileStoragePath string) *IPLDServiceClient {
	return &IPLDServiceClient{
		client: client,
		FileStoragePath: FileStoragePath,
	}
}

func (c *IPLDServiceClient) UploadFile(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	res, err := c.client.UploadFile(context.Background(), &pb.FileUploadRequest{FileData: fileData})
	if err != nil {
		return "", err
	}

	return res.Cid, nil
}

func (c *IPLDServiceClient) GetFileByCID(cid string) []byte {
	res, err := c.client.GetFileByCID(context.Background(), &pb.FileRequest{Cid: cid})
	if err != nil {
		return nil
	}

	err = os.WriteFile(c.FileStoragePath + "/" + cid, res.FileData, 0644)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return res.FileData
}