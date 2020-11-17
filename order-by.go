package visitypes

type Order string

const (
	OrderAsc  = "ASC"
	OrderDesc = "DESC"
)

type OrderBy struct {
	field string
	order Order
}

func NewOrderBy(field string, order Order) *OrderBy {
	return &OrderBy{field: field, order: order}
}

func (ob *OrderBy) GetField() string {
	return ob.field
}

func (ob *OrderBy) GetOrder() Order {
	return ob.order
}
