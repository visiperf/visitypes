package visitypes

type Pagination struct {
	start int
	limit int
}

func NewPagination(start, limit int) *Pagination {
	return &Pagination{start: start, limit: limit}
}

func (p *Pagination) GetStart() int {
	return p.start
}

func (p *Pagination) GetLimit() int {
	return p.limit
}
