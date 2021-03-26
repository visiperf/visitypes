package visitypes

type Order int

const (
	OrderAsc Order = iota
	OrderDesc
)

type OrderBy struct {
	field string
	order Order
}

func NewOrderBy(field string, order Order) *OrderBy {
	return &OrderBy{field, order}
}

func (ob OrderBy) Field() string {
	return ob.field
}

func (ob OrderBy) Order() Order {
	return ob.order
}
