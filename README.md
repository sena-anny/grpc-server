## 事前準備
```zsh
brew install protobuf
export GO111MODULE=on
go get google.golang.org/protobuf/cmd/protoc-gen-go \
      google.golang.org/grpc/cmd/protoc-gen-go-grpc
export PATH="$PATH:$(go env GOPATH)/bin"
```

## protoファイルからGoファイル作成
```zsh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/helloworld.proto
```

## 
```zsh
go mod init kj-tech.net/module


```
