#!/bin/bash

PROTO_FILE=fizzbuzz.proto
SRC_DIR=.
DST_DIR=.

protoc --proto_path=$SRC_DIR --go_out=plugins=grpc:$DST_DIR $PROTO_FILE
