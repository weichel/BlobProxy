package main

import (
	"context"
	"testing"
	"time"
	"github.com/golang/mock/gomock"
	"fmt"
	"github.com/golang/protobuf/proto"
	"crypto/sha256"
	pb "github.com/weichel/BlobProxy/proto"
	pbmocks "github.com/weichel/BlobProxy/mocks"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}
func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}
func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

// Protocol Buffer Mock Testing On ReadBlob
func TestReadBlob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := pbmocks.NewMockBlobProxyClient(ctrl)

 	req := &pb.ReadBlobRequest{Key: "unit_test"}

	mockClient.EXPECT().ReadBlob(
			gomock.Any(),
			&rpcMsg{msg: req},
	).Return(&pb.ReadBlobResponse{Data: []byte("Mock BlobProxy Client Test")}, nil)

	testReadBlob(t, mockClient)
}
// testReadBlob helper
func testReadBlob(t *testing.T, client pb.BlobProxyClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.ReadBlob(ctx, &pb.ReadBlobRequest{Key: "unit_test"})
	if err != nil || string(r.Data) != "Mock BlobProxy Client Test" {
		t.Errorf("mocking failed")
	}
}

// Server Data Testing
func TestServer(t *testing.T) {
	s := server{}

	/*
		Test cases, set to look for a specific Sha256 hash string
	   This structure extended much further to include
		 		user input data sanitization
				fail cases for other calls in server such as http.Get and net.listen, etc
	*/
	tests := []struct{
	        key string
	        want string
	    } {
					{
							key: "lorum.txt",
							want: "0b2ea3e6d9c9f48d7b1bbebbd0ff64c0e1d60c05cf6415d5b32e4df7790bc455",
					},
					{
							key: "dog.jpg",
							want: "a770895136a1970647ed3bf74b1ff485599590b89b4a9852c26c7b0bb057f36e",
					},
					{
							key: "cat.jpg",
							want: "9d5158a0ccef225168947493a6080f89d192021f49160ea142374d707b26c309",
					},
	    }

	// Loop over test cases
	for _, tt := range tests {
		// Launch azure blob request
    req := &pb.ReadBlobRequest{Key: tt.key}
    resp, err := s.ReadBlob(context.Background(), req)
    if err != nil {
        t.Errorf("testServer(%v) got unexpected error",tt.key)
    }

		// Check Sha256 string to verify azure contents
    if fmt.Sprintf("%x", sha256.Sum256(resp.Data)) != tt.want {
        t.Errorf("Sha256(ReadBlob(%v))=%v, wanted %v", tt.key, sha256.Sum256(resp.Data), tt.want)
    }
	}

}
