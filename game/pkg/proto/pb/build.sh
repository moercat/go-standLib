#!/bin/sh
SRC_DIR=.
DST_DIR=.
rm -rf ../api/*.go
protoc --proto_path=. --go_out=plugins=grpc:../api *.proto