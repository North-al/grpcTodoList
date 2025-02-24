CreateUserProto:
	protoc -I=TodoListUser/api/proto \
		--go_out=./gen \
		--go-grpc_out=./gen \
		./TodoListUser/api/proto/user.proto