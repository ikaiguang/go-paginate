# api

generate protobuf

```shell
protoc -I. --go_out=. --go_opt=paths=source_relative ./*.proto
```