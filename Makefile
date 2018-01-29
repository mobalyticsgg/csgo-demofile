
setup:
	go get github.com/gogo/protobuf/protoc-gen-gogofaster

proto_build:
	protoc -I=proto --gogofaster_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:proto/. proto/*.proto