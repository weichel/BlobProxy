# gRPC BlobProxy Server

This solution contains the gRPC Server implemented in Go. 

There is also a node.js package that can be installed run a node server and client (but is not referenced below).

## How To Run

### BlobProxy Server on Docker Container

The Dockerfile in the /go directory will create a golang image with:
* Environment variables set appropriately for Azure
* Latest version of Go, a copy of the BlobProxy source code, and an installation of the BlobProxy
* Server port exposed and set to launch on container startup

1. Create the image from Dockerfile: 'docker install -t server_image'
2. Next, start the server container: 'docker run -it -p 1337:1337 --name Server server_image'
3. Run the command: "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' Server"
   (We will use this to connect the client container)


### BlobProxy Client on Docker Container

Similar to above, but start with the Dockerfile in /go/client

1. Create the image from Dockerfile: 'docker install -t client_image'
2. Next, start the server container: 'docker run -it -P client_image <blob.key> <Address from above:port>'

You should see the results of the blob data printed to the screen. 

### Node.js Client
The node.js client can quickly be used to test the server by running: 'node client <blob.key>'


## Running Test Cases

Go makes it very easy to compile and test our project. 
In the github.com/weichel/BlobProxy directory, run the command: "go test -v ."

