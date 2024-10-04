protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --endpoint_out=services=true:. ./example/proto/_.proto
protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --endpoint_out=types=true:. ./example/proto/_.proto
protoc --plugin=protoc-gen-gbtemplate=./protoc-gen-gbtemplate --endpoint_out=methods=true:. ./example/proto/\*.proto
