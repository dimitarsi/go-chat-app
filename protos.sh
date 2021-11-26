MODULE_NAME="github.com/dimitarsi/go-chat-app"

protoc --proto_path=. ./proto/*.proto \
    --go_out=. \
    --go-grpc_out=. \
    --go_opt=module=$MODULE_NAME \
    --go-grpc_opt=module=$MODULE_NAME