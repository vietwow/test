protoc -I pb/ \
-I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
pb/user.proto \
--go_out=plugins=grpc:pb
