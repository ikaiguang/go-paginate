# paginate

paginate for golang

## 请求协议

api protobuf

```protobuf

syntax = "proto3";

// PageRequest 分页请求
message PageRequest {
  // page 第几页；默认第一页
  uint32 page = 100;
  // page_size 一页显示的数据条数
  uint32 page_size = 101;

  // order_by_array 排序；例：["id,desc", "created_at,asc"]
  repeated string order_by_array = 200;
}

// PageInfo 分页信息
message PageInfo {

  // total_number 总条数
  int64 total_number = 100;
  // total_page 总页数
  int64 total_page = 101;
  // page 第几页
  uint32 page = 102;
  // page_size 一页显示的数据条数
  uint32 page_size = 103;
}

```

## 分页选项

[分页接口与定义](./paginate_reference.go)

```go
package paginate

import paginatepb "github.com/ikaiguang/go-paginate/api"

// Interface page interface
type Interface interface {
    // 分页选项
    MakePageOptions(*paginatepb.PageRequest) *Options
    
    // 默认请求
    DefaultPageRequest() *paginatepb.PageRequest
    // 解析分页请求
    ParsePageRequest(*paginatepb.PageRequest) *paginatepb.PageRequest
    // 获取分页方向
    ParseDirection(string) Direction
}

// Direction .
type Direction string

// Options .
type Options struct {
	Request *paginatepb.PageRequest

	Where  []*Where
	Order  []*OrderBy
	Limit  uint32
	Offset uint32
}

// Where 分页条件；例：where id = ?(where id = 1)
type Where struct {
	// Column 字段
	Column string
	// Symbol 条件
	Symbol string
	// Placeholder 占位符
	Placeholder string
	// Data 数据
	Data interface{}
}

// OrderBy 排序
type OrderBy struct {
	// Column 字段
	Column string
	// Direction 排序方向
	Direction Direction
}
```
