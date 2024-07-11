# gRPC Protobuf Ping VM
The purpose of this project is to demonstrate a gRPC-based client-server communication setup using protobuf. The project involves setting up a server on an EC2 instance and a client that communicates with this server.

## Requirements
* Go (latest version) on your VM and local machine
* EC2 instance (& an open port 50051 on the security dashboard)

## Setup

### 1. Push the entire project directory to your VM
From the parent directory of your project, run:
```
scp -i <pem-key> -r grpc-protobuf-ping-vm user@VM_IP_ADDR:~
```

### 2. Run the server on the VM
SSH into your VM, navigate to the project directory, and run:
```
cd ~/grpc-protobuf-ping-vm/server
go build -o server_exec main.go
./server_exec
```

### 3. Run the client locally
Navigate to the project directory on your local machine and run:
```
go run client/main.go
```

### Expected Result
```
<timestamp> Response: Pong: Hello, Server!
```