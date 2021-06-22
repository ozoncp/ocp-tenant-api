run:
	go run cmd/ocp-tenant-api/main.go

lint:
	golint ./...

test:
	go test -v ./...

.PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-tenant-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-tenant-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-tenant-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-tenant-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-tenant-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-tenant-api/ocp-tenant-api.proto
		mv pkg/ocp-tenant-api/github.com/ozoncp/ocp-tenant-api/pkg/ocp-tenant-api/* pkg/ocp-tenant-api/
		rm -rf pkg/ocp-tenant-api/gihtub.com
		mkdir -p cmd/ocp-tenant-api

PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-tenant-api cmd/ocp-tenant-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-tenant-api
		cp api/ocp-tenant-api/ocp-tenant-api.proto vendor.protogen/api/ocp-tenant-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc