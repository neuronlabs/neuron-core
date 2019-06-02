package query

import (
	"github.com/neuronlabs/neuron/internal"
	"github.com/neuronlabs/neuron/internal/query/paginations"
	"github.com/neuronlabs/neuron/log"
	"net/url"
	"strconv"
)

// PaginationType defines the pagination type
type PaginationType int

const (
	// TpLimitOffset is the pagination type that defines limit or (and) offset
	TpLimitOffset PaginationType = iota

	// TpPage is the pagination type that uses page type pagination i.e. page=1 page size = 10
	TpPage
)

// newPagination creates new pagination for the given query
// The arguments 'a' and 'b' means as follows:
// - 'limit' and 'offset' for the type 'TpLimitOffset'
// - 'pageNumber' and 'pageSize' for the type 'TpPage'
func newPagination(a, b int, tp PaginationType) *Pagination {
	p := &paginations.Pagination{}
	switch tp {
	case TpLimitOffset:
		p.SetValue(a, paginations.ParamLimit)
		p.SetValue(b, paginations.ParamOffset)
		p.SetType(paginations.TpOffset)
	case TpPage:
		p.SetValue(a, paginations.ParamNumber)
		p.SetValue(b, paginations.ParamSize)
		p.SetType(paginations.TpPage)
	default:
		log.Debugf("Pagination type not found.")
		// not implemented yet
	}
	return (*Pagination)(p)
}

// Pagination defines the query limits and offsets.
// It defines the maximum size (Limit) as well as an offset at which
// the query should start.
type Pagination paginations.Pagination

// SetLimit sets the limit for the pagination
func (p *Pagination) SetLimit(limit int) {
	(*paginations.Pagination)(p).SetValue(limit, paginations.ParamLimit)
}

// SetOffset sets the offset for the pagination
func (p *Pagination) SetOffset(offset int) {
	(*paginations.Pagination)(p).SetValue(offset, paginations.ParamOffset)
}

// SetPageNumber sets the page number for the pagination
func (p *Pagination) SetPageNumber(pageNumber int) {
	(*paginations.Pagination)(p).SetValue(pageNumber, paginations.ParamNumber)
}

// SetPageSize sets the page number for the pagination
func (p *Pagination) SetPageSize(pageSize int) {
	(*paginations.Pagination)(p).SetValue(pageSize, paginations.ParamSize)
}

// PagedPagination creates new paged type Pagination
func PagedPagination(number, size int) *Pagination {
	return newPaged(number, size)
}

func newPaged(number, size int) *Pagination {
	return (*Pagination)(paginations.NewPaged(number, size))
}

// LimitOffsetPagination createw new limit, offset type pagination
func LimitOffsetPagination(limit, offset int) *Pagination {
	return newLimitOffset(limit, offset)
}

func newLimitOffset(limit, offset int) *Pagination {
	return (*Pagination)(paginations.NewLimitOffset(limit, offset))
}

// GetLimitOffset gets the limit and offset from the current pagination
func (p *Pagination) GetLimitOffset() (int, int) {
	return (*paginations.Pagination)(p).GetLimitOffset()
}

// Check checks if the pagination is well formed
func (p *Pagination) Check() error {
	return (*paginations.Pagination)(p).Check()
}

// String implements Stringer interface
func (p *Pagination) String() string {
	return (*paginations.Pagination)(p).String()
}

// Type returns pagination type
func (p *Pagination) Type() PaginationType {
	return PaginationType((*paginations.Pagination)(p).Type())
}

// FormatQuery formats the pagination for the url query.
func (p *Pagination) FormatQuery(q ...url.Values) url.Values {

	var query url.Values
	if len(q) != 0 {
		query = q[0]
	}

	if query == nil {
		query = url.Values{}
	}

	var k, v string

	switch p.Type() {
	case TpLimitOffset:
		limit, offset := p.GetLimitOffset()
		if limit != 0 {
			k = internal.QueryParamPageLimit
			v = strconv.Itoa(limit)
			query.Set(k, v)
		}
		if offset != 0 {
			k = internal.QueryParamPageOffset
			v = strconv.Itoa(offset)
			query.Set(k, v)
		}
	case TpPage:
		number, size := (*paginations.Pagination)(p).GetNumberSize()
		if number != 0 {
			k = internal.QueryParamPageNumber
			v = strconv.Itoa(number)
			query.Set(k, v)
		}

		if size != 0 {
			k = internal.QueryParamPageSize
			v = strconv.Itoa(size)
			query.Set(k, v)
		}

	default:
		log.Debugf("Pagination with invalid pagination type: '%s'", p.Type())
	}
	return query
}
