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

// ToDirection .
func ToDirection(val string) Direction {
	return Direction(val)
}

const (
	// 分页
	DefaultPage     = 1
	DefaultPageSize = 20

	// 排序
	DefaultDirectionAsc  Direction = "ASC"
	DefaultDirectionDesc Direction = "DESC"
)

// Value .
func (d Direction) Value() string {
	return string(d)
}

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
