# neonyx-file-ipld
## IPLD File Service
The IPLD File Service is a decentralized file storage service that leverages the InterPlanetary Linked Data (IPLD) structure to store and retrieve files. It is built using Go and utilizes gRPC for communication between the client and server.

### APIs
Here, we provide a description of the APIs that facilitate file uploads and downloads:

### Server APIs
**UploadFile(ctx context.Context, req pb.FileUploadRequest) (pb.FileUploadResponse, error)

Uploads a file to the server. It accepts a FileUploadRequest that contains the file data and the file name, and returns a FileUploadResponse containing the CID (Content Identifier) which can be used later to retrieve the file.

**GetFileByCID(ctx context.Context, req pb.FileRequest) (pb.FileResponse, error)

Retrieves a file from the server using its CID. It accepts a FileRequest that contains the CID, and returns a FileResponse containing the file data and the file name.

### Client APIs
UploadFile(filePath string) (string, error)

Facilitates the upload of a file from the specified file path to the server. It returns the CID generated upon the successful upload of the file.

GetFileByCID(cid, outputDirectory string) error

Facilitates the retrieval of a file using its CID and saves it to the specified output directory, preserving the original file name.

### Usage
To interact with the IPLD File Service, you will primarily work with the client APIs, which abstract the gRPC communications for easy usage in Go programs.
