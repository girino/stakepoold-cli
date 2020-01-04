# stakepoold-cli
Command line tools for accessing stakepoold remotely

Not yet funcional

## install

prerequisites: 
```
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ export PATH=$PATH:$GOPATH/bin
$ wget https://raw.githubusercontent.com/decred/dcrstakepool/v1.2.0/backend/stakepoold/rpc/api.proto
$ mkdir stakepoolrpc
$ protoc api.proto --go_out=plugins=grpc:stakepoolrpc
```

## Build

`git clone https://github.com/girino/stakepoold-cli`
`go build`

## Run

`./stakepoold-cli -h host -p port -c /path/to/cert ping`
