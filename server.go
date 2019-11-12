package main

import (
	"context"
	"log"
	//"fmt"
	"net"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"google.golang.org/grpc"
	pb "github.com/weichel/BlobProxy/proto"
)

type server struct {
	pb.UnimplementedBlobProxyServer
}

// Create the base uri needed for the Azure Blob Storeage REST call made in the ReadBlob RPC
var base_uri = strings.Join([]string{"https://",os.Getenv("BLOB_ACCOUNT"),".blob.core.windows.net/",os.Getenv("BLOB_CONTAINER"),"/"},"")
const port = ":1337"

/*******************************************************************
  ReadBlob RPC defined in ./barracuda.proto
        in: ReadBlobRequest.key - blob key to request from the Azure Storage Account/Container
								in the environment variables BLOB_ACCOUNT and BLOB_CONTAINER

				out: ReadBlobResponse.data - Body of the http response containing the blob data


  The ReadBlob RPC is centered around an HTTP Get request to the public Azure Blob Storage
	account stored in the environment variables. The body of the response is read and sent
	on to the client.
*******************************************************************/
func (s *server) ReadBlob(ctx context.Context, req *pb.ReadBlobRequest) (*pb.ReadBlobResponse, error) {

	// Run a Get request to Azure with the incoming blob key
	resp, err := http.Get(strings.Join([]string{base_uri,req.Key},""))
	if err != nil {
		log.Fatalf("failed get request: %v", err)
	}
	defer resp.Body.Close()

	// Read the body of the http response containing the blob data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed response read: %v", err)
	}

	// Return the body of the response to the client
	return &pb.ReadBlobResponse{Data: body}, nil
}

func main() {
	// Start up a tcp socket for the gRPC Service
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create the gRPC server, register BlobProxy proto, and launch the service
	s := grpc.NewServer()
	pb.RegisterBlobProxyServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
