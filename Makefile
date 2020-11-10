run-server:
	go run cmd/server/main.go
run-client:
	go run cmd/client/main.go
run-basic:
	go run cmd/basic/main.go
run-basic-client:
	go run cmd/basic-client/main.go

tidy:
	go mod tidy
	go mod vendor
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

gen:
	protoc -I . \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		api/*/*.proto

	protoc -I . \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=api/user/ \
		--openapiv2_out=api/user/swagger \
		api/user/*.proto

clean-proto:
	rm api/*/*.go
clean-swagger:
	rm -r api/*/swagger/*