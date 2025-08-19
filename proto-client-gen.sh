#!/bin/bash
rm -rf ./client/src/proto
# protoc -I=. ./server/proto/*.proto \
#   --js_out=import_style=commonjs,binary:./client/src \
#   --grpc-web_out=import_style=typescript,mode=grpcwebtext:./client/src

protoc -I=. ./server/proto/*.proto \
  --ts_out=./client/src


mv -f ./client/src/server/proto  ./client/src

rm -rf ./client/src/server