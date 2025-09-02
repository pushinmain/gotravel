PROTO_PATH=pkg/protos/ 

proto-gen:
	@protoc -I proto ${PROTO_PATH}/proto/sso/sso.proto --go_out=${PROTO_PATH}/gen/go --go_opt=paths=source_relative --go-grpc_out=${PROTO_PATH}/gen/go --go-grpc_opt=paths=source_relative 

.PHONY: proto-gen