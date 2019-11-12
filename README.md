# gRPC BlobProxy Server

This solution contains the gRPC Server implemented in Go. 

There is also a node.js package that can be installed run a node server and client (but is not referenced below).

## How To Run

### BlobProxy Server on Docker Container

The Dockerfile in the /go directory will create a golang image with:
* Environment variables set appropriately for Azure
* Latest version of Go, a copy of the BlobProxy source code, and an installation of the BlobProxy
* Server port exposed and set to launch on container startup

Run the command: 'docker install -t server_image'
Next, start the server container: 'docker run -it -p 1337:1337 server_image'


### [Node.js Developer Tools](https://github.com/docker/labs/blob/master/developer-tools/nodejs-debugging/README.md) including:
+ Visual Studio Code

## Programming languages
This is a more comprehensive section detailing how to set-up and optimize your experience using Docker with particular programming languages.

+ [Java](java/)
+ [Node.js](nodejs/porting/)
