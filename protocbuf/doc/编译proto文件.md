go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/proto.proto