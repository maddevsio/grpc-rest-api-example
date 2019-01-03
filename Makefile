TARGET=helloworld

all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) main.go

proto:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--go_out=plugins=grpc:. \
		pb/hello.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--grpc-gateway_out=logtostderr=true:. \
		pb/hello.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--swagger_out=logtostderr=true:. \
		pb/hello.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/lyft/protoc-gen-validate \
		--validate_out="lang=go:." \
		pb/hello.proto
