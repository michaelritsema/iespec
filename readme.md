# Build protobufs on OSX
# Change -I=/usr/local/Cellar/protobuf/2.6.0/include to where ever it is installed on your system

protoc --go_out=plugins=grpc,import_prefix=,import_path=protomsg:$(pwd)/protomsg  -I=/usr/local/Cellar/protobuf/2.6.0/include -I $(pwd)/protomsg/sources $(pwd)/protomsg/sources/*.proto


extract protobuf-descriptor in your $GO_ROOT folder

building requirements:
yum install libpcap-devel
git version  > 1.7.2 (use repoforge for centos)
