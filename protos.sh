#!/bin/bash
protoc --proto_path=. ./proto/*.proto \
    --go_out=. \
    --go-grpc_out=. \
    --go_opt=module=github.com/dimitarsi/go-chat-app \
    --go-grpc_opt=module=github.com/dimitarsi/go-chat-app