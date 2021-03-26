package visitypes

type Pagination struct {
	start int
	limit int
}

func NewPagination(start, limit int) *Pagination {
	return &Pagination{start, limit}
}

func (p Pagination) Start() int {
	return p.start
}

func (p Pagination) Limit() int {
	return p.limit
}
