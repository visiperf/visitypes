package visitypes

type Operator int

const (
	OperatorIn Operator = iota
	OperatorEqual
	OperatorLike
	OperatorIsNull
	OperatorLessThan
	OperatorGreaterThan
	OperatorBetween
)

type Predicate struct {
	field    string
	operator Operator
	values   []interface{}
}

func NewPredicate(field string, operator Operator, values []interface{}) *Predicate {
	return &Predicate{field, operator, values}
}

func (p Predicate) Field() string {
	return p.field
}

func (p Predicate) Operator() Operator {
	return p.operator
}

func (p Predicate) Values() []interface{} {
	return p.values
}
