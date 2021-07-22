package paginate

import (
	paginatepb "github.com/ikaiguang/go-paginate/api"
)

//var _ Interface = (*Paginate)(nil)

// Paginate 实现Interface
type Paginate struct{}

// MakePageOptions .
func (p *Paginate) MakePageOptions(req *paginatepb.PageRequest) *Options {
	req = p.ParsePageRequest(req)

	return p.makePageOptions(req)
}

// DefaultPageRequest .
func (p *Paginate) DefaultPageRequest() *paginatepb.PageRequest {
	return &paginatepb.PageRequest{
		Page:     DefaultPage,
		PageSize: DefaultPageSize,

		OrderByArray: []string{},
	}
}

// ParsePageRequest .
func (p *Paginate) ParsePageRequest(req *paginatepb.PageRequest) *paginatepb.PageRequest {
	if req == nil {
		return p.DefaultPageRequest()
	}
	return p.parsePageRequest(req)
}

// makePageOptions .
func (p *Paginate) makePageOptions(req *paginatepb.PageRequest) (opts *Options) {
	opts = &Options{
		Request: req,

		Where:  []*Where{},
		Order:  []*OrderBy{},
		Limit:  req.PageSize,
		Offset: req.PageSize * (req.Page - 1),
	}
	return opts
}
