proto:
	protoc -I ./crud/server/api --go_out=plugins=grpc:./crud/server/api crud/server/api/*.proto

