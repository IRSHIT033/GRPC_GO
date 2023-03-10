
# GRPC_GO

![alt text](https://cdn.hashnode.com/res/hashnode/image/upload/v1676286888080/32e693e3-1d2f-40b2-a427-408c91ecd2b0.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress,format&format=webp)


## How to start the application  
 
 1. The first step is to install the Go plugins for the protocol compiler:

  ```
      go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
      go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.22
  ```

  2. Update your PATH so that the protoc compiler can find the plugins:

  ```
      export PATH="$PATH:$(go env GOPATH)/bin"
  ```
  3. Then run this command to install all the needed dependency
  ```
      go mod tidy
  ```
  4. Next build and run the client and server to start application
     
     - open a terminal to run server
     
  ``` 
      cd server
      go build 
      ./server
  ```
  
     - open another terminal to run client
     
  ``` 
       cd client
       go build 
       ./client
  ```
 
