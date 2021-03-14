

protoc -I ./helloworld/ --go_out=plugins=grpc:helloworld helloworld/helloworld.proto
protoc -I helloworld --grpc-gateway_out=logtostderr=true:helloworld/ helloworld/helloworld.proto
protoc -I helloworld/ --swagger_out=logtostderr=true:helloworld/ helloworld/helloworld.proto

protoc -I ./echo/ --go_out=plugins=grpc:echo echo/echo_service.proto
protoc -I echo --grpc-gateway_out=logtostderr=true:echo/ echo/echo_service.proto
protoc -I echo/ --swagger_out=logtostderr=true:echo/ echo/echo_service.proto