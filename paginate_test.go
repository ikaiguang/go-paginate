package paginate

import (
	paginatepb "github.com/ikaiguang/go-paginate/api"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// var
var (
	pageHandler Interface
)

func TestMain(m *testing.M) {
	pageHandler = &Paginate{}

	code := m.Run()

	os.Exit(code)
}

// go test -v -count=1 -test.run=TestPaginate_MakePageOptions
func TestPaginate_MakePageOptions(t *testing.T) {
	defaultRequest := pageHandler.DefaultPageRequest()
	pageRequestFor10 := &paginatepb.PageRequest{
		Page:     10,
		PageSize: 10,
	}

	tests := []struct {
		name  string
		given *paginatepb.PageRequest
		want  *Options
	}{
		{
			name:  "#准备分页选项#defaut",
			given: defaultRequest,
			want: &Options{
				Request: defaultRequest,

				Where:  []*Where{},
				Order:  []*OrderBy{},
				Limit:  defaultRequest.PageSize,
				Offset: 0,
			},
		},
		{
			name:  "#准备分页选项#每页10条之第10页",
			given: pageRequestFor10,
			want: &Options{
				Request: pageRequestFor10,

				Where:  []*Where{},
				Order:  []*OrderBy{},
				Limit:  pageRequestFor10.PageSize,
				Offset: (pageRequestFor10.Page - 1) * pageRequestFor10.PageSize,
			},
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := pageHandler.MakePageOptions(param.given)
			assert.Equal(t, param.want.Limit, got.Limit, "Limit")
			assert.Equal(t, param.want.Offset, got.Offset, "Offset")
		})
	}
}

// go test -v -count=1 -test.run=TestPaginate_ParsePageRequest
func TestPaginate_ParsePageRequest(t *testing.T) {
	tests := []struct {
		name  string
		given *paginatepb.PageRequest
		want  *paginatepb.PageRequest
	}{
		{
			name:  "#解析分页请求#nil",
			given: nil,
			want: &paginatepb.PageRequest{
				Page:         DefaultPage,
				PageSize:     DefaultPageSize,
				OrderByArray: []string{},
			},
		},
		{
			name: "#解析分页请求#default",
			given: &paginatepb.PageRequest{
				Page:         DefaultPage,
				PageSize:     DefaultPageSize,
				OrderByArray: []string{},
			},
			want: &paginatepb.PageRequest{
				Page:         DefaultPage,
				PageSize:     DefaultPageSize,
				OrderByArray: []string{},
			},
		},
		{
			name: "#解析分页请求#zero",
			given: &paginatepb.PageRequest{
				Page:         0,
				PageSize:     0,
				OrderByArray: []string{},
			},
			want: &paginatepb.PageRequest{
				Page:         DefaultPage,
				PageSize:     DefaultPageSize,
				OrderByArray: []string{},
			},
		},
		{
			name: "#解析分页请求#custom",
			given: &paginatepb.PageRequest{
				Page:         2,
				PageSize:     300,
				OrderByArray: []string{},
			},
			want: &paginatepb.PageRequest{
				Page:         2,
				PageSize:     300,
				OrderByArray: []string{},
			},
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := pageHandler.ParsePageRequest(param.given)
			assert.Equal(t, param.want.Page, got.Page, "Page")
			assert.Equal(t, param.want.PageSize, got.PageSize, "PageSize")
			assert.Equal(t, len(param.want.OrderByArray), len(got.OrderByArray), "OrderByArray")
		})
	}
}

// go test -v -count=1 -test.run=TestPaginate_ParseDirection
func TestPaginate_ParseDirection(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  Direction
	}{
		{
			name:  "#解析分页排序方向#unknown",
			given: "unknown",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#desc",
			given: "desc",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#DESC",
			given: "DESC",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#DeSC",
			given: "DeSC",
			want:  DefaultDirectionDesc,
		},
		{
			name:  "#解析分页排序方向#asc",
			given: "asc",
			want:  DefaultDirectionAsc,
		},
		{
			name:  "#解析分页排序方向#ASC",
			given: "ASC",
			want:  DefaultDirectionAsc,
		},
		{
			name:  "#解析分页排序方向#AsC",
			given: "AsC",
			want:  DefaultDirectionAsc,
		},
	}
	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := pageHandler.ParseDirection(param.given)
			assert.Equal(t, param.want, got, "Direction")
		})
	}
}

// go test -v -count=1 -test.run=TestPaginate_DefaultPageRequest
func TestPaginate_DefaultPageRequest(t *testing.T) {
	tests := []struct {
		name string
		want *paginatepb.PageRequest
	}{
		{
			name: "#默认的分页请求",
			want: &paginatepb.PageRequest{
				Page:         DefaultPage,
				PageSize:     DefaultPageSize,
				OrderByArray: []string{},
			},
		},
	}

	for _, param := range tests {
		t.Run(param.name, func(t *testing.T) {
			got := pageHandler.DefaultPageRequest()
			assert.Equal(t, param.want.Page, got.Page, "Page")
			assert.Equal(t, param.want.PageSize, got.PageSize, "PageSize")
			assert.Equal(t, len(param.want.OrderByArray), len(got.OrderByArray), "OrderByArray")
		})
	}
}
