set PROTO_FILE=fizzbuzz.proto
set SRC_DIR=.
set DST_DIR=.

protoc --proto_path=%SRC_DIR% --go_out=plugins=grpc:%DST_DIR% %PROTO_FILE%
