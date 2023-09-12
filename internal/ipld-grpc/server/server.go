package server

import (
	"context"

	pb "github.com/KadirbekSharau/file-to-ipld/internal/ipld-grpc/api/ipld"
	fileipld "github.com/KadirbekSharau/file-to-ipld/pkg/file/file-ipld"
)

type IPLDServiceServer struct {
	pb.UnimplementedIPLDServiceServer
    fileManager fileipld.IPLDFileManager
}

func NewIPLDServiceServer(fileManager fileipld.IPLDFileManager) *IPLDServiceServer {
	return &IPLDServiceServer{
        fileManager: fileManager,
    }
}

func (s *IPLDServiceServer) UploadFile(ctx context.Context, req *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	cid, err := s.fileManager.UploadFileToDisk(req.FileData)

	if err != nil {
		return nil, err
	}

	return &pb.FileUploadResponse{Cid: cid.String()}, nil
}

func (s *IPLDServiceServer) GetFileByCID(ctx context.Context, req *pb.FileRequest) (*pb.FileResponse, error) {
    fileData, err := s.fileManager.GetFileFromDiskByCID(req.Cid)
	if err != nil {
		return nil, err
	}

	return &pb.FileResponse{FileData: fileData}, nil
}
