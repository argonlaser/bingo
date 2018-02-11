echo "Generate protocol buffer file"
protoc -I gameShipRpc/  gameShipRpc/gameShipRpc.proto  --go_out=plugins=grpc:gameShipRpc
