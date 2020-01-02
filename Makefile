proto:
	protoc -I ./server/api --go_out=plugins=grpc:./server/api server/api/*.proto

