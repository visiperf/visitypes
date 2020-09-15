package visitypes

type Operator string

const (
	OperatorIn          Operator = "IN"
	OperatorEqual       Operator = "EQUALS"
	OperatorLike        Operator = "LIKE"
	OperatorIsNull      Operator = "IS NULL"
	OperatorLessThan    Operator = "LESS THAN"
	OperatorGreaterThan Operator = "GREATER THAN"
)

type Predicate struct {
	field    string
	operator Operator
	values   []interface{}
}

func NewPredicate(field string, operator Operator, values []interface{}) *Predicate {
	return &Predicate{field: field, operator: operator, values: values}
}

func (p *Predicate) GetField() string {
	return p.field
}

func (p *Predicate) GetValues() []interface{} {
	return p.values
}

func (p *Predicate) IsOperator(operator string) bool {
	return operator == string(p.operator)
}
