package paginate

import (
	paginatepb "github.com/ikaiguang/go-paginate/api"
	"google.golang.org/protobuf/proto"
	"strings"
)

// ParseDirection .
func (p *Paginate) ParseDirection(direction string) Direction {
	switch od := ToDirection(strings.ToUpper(direction)); od {
	case DefaultDirectionDesc, DefaultDirectionAsc:
		return od
	default:
		return DefaultDirectionDesc
	}
}

// parsePageRequest .
func (p *Paginate) parsePageRequest(req *paginatepb.PageRequest) *paginatepb.PageRequest {
	req = proto.Clone(req).(*paginatepb.PageRequest)

	req.Page = p.parsePage(req.Page)
	req.PageSize = p.parsePageSize(req.PageSize)
	req.OrderByArray = []string{}

	return req
}

// parsePage .
func (p *Paginate) parsePage(page uint32) uint32 {
	if page < 1 {
		page = DefaultPage
	}
	return page
}

// parsePageSize
func (p *Paginate) parsePageSize(size uint32) uint32 {
	if size < 1 {
		size = DefaultPageSize
	}
	return size
}
