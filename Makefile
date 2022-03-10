gen:
	protoc -I=./messages --go_out=plugins=grpc:. ./messages/*.proto

get:
	go mod tidy: