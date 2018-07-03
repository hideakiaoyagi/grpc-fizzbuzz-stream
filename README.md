# gRPC-FizzBuzz-stream

A sample of gRPC programming in Golang.

This repository is a version that supported the following repository to server-side-streaming.  
https://github.com/hideakiaoyagi/grpc-fizzbuzz


## About gRPC

https://grpc.io/  
https://github.com/grpc/  


## Instructions

### Instructions for Golang

#### 1. Install gRPC, Protocol Buffers, and more tools.  
see: https://grpc.io/docs/quickstart/go.html#prerequisites

#### 2. Compile the proto file into a Gplang-stub file.
```
$ cd proto
$ ./generate_golang_code.sh
$ ls
fizzbuzz.pb.go
```

#### 3. Run server.
```
$ cd server_golang
$ go run fizzbuzz_server.go
```

#### 4. Start another terminal, and run client.
```
$ cd client_golang
$ go run fizzbuzz_client.go
1
2
Fizz
4
Buzz
...
```


### Instructions for C#

Now, not implemented.
