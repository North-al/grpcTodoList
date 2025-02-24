CreateUserProto:
	protoc -I=app/user/api/proto \
		--go_out=./gen \
		--go-grpc_out=./gen \
		./app/user/api/proto/user.proto