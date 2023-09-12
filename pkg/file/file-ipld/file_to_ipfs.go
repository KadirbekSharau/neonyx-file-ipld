package fileipld

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	ipfscid "github.com/ipfs/go-cid"
	"github.com/ipfs/go-merkledag"
)

type IPLDFileManager interface {
	UploadFileToDisk(fileData []byte)  (ipfscid.Cid, error)
	GetFileFromDiskByCID(cidStr string) ([]byte, error)
}

type fileManager struct {
	IPLDStoragePath string
}

func NewIPLDFileManager(IPLDStoragePath string) IPLDFileManager {
	return &fileManager{
		IPLDStoragePath: IPLDStoragePath,
	}
}

// Deprecate
func (p *fileManager) ReceiveFile(filePath string) ([]byte, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

// Deprecate
func (p *fileManager) CreateNodeIPLD(encryptedData []byte) ([]byte, error) {
	node := merkledag.NodeWithData(bytes.NewBuffer(encryptedData).Bytes())
	if _, err := node.EncodeProtobuf(false); err != nil {
		fmt.Println("Failed to encode to protobuf:", err)
		return nil, err
	}
	serializedData, err := json.Marshal(node)
	if err != nil {
		fmt.Println("Failed to serialize node:", err)
		return nil, err
	}
	return serializedData, nil
}

// Deprecate
func (p *fileManager) UploadToDisk(ipldData []byte) error {
	err := os.WriteFile("ipld_structure", ipldData, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write IPLD structure to disk:", err)
		return err
	}
	return nil
}

func (p *fileManager) UploadFileToDisk(fileData []byte) (ipfscid.Cid, error) {
	encryptedData := encryptFile(fileData)

	node := merkledag.NodeWithData(bytes.NewBuffer(encryptedData).Bytes())

	cid := node.Cid()

	serializedData, err := json.Marshal(node)
	if err != nil {
		fmt.Println("Failed to serialize node:", err)
		return ipfscid.Undef, nil
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s.json", p.IPLDStoragePath, cid.String()), serializedData, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to write IPLD structure to disk:", err)
		return ipfscid.Undef, nil
	}

	fmt.Println("Successfully converted file to IPLD and saved to disk.")
	return cid, nil
}

func (p *fileManager) GetFileFromDiskByCID(cidStr string) ([]byte, error) {
		serializedData, err := os.ReadFile(fmt.Sprintf("%s/%s.json", p.IPLDStoragePath, cidStr))
		if err != nil {
			return nil, fmt.Errorf("failed to read IPLD data from disk: %v", err)
		}
	
		var node merkledag.ProtoNode
		if err := json.Unmarshal(serializedData, &node); err != nil {
			return nil, fmt.Errorf("failed to deserialize IPLD node: %v", err)
		}
	
		fileData := node.Data()
		decryptedData := decryptFile(fileData)
	
		return decryptedData, nil
}

// TODO()
func encryptFile(data []byte) []byte {
	return data
}

// TODO()
func decryptFile(data []byte) []byte {
	return data
}
