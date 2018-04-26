package query

type Operator string

const (
	GROUP Operator = "GROUP"
	AND   Operator = "AND"
	OR    Operator = "OR"
	NOT   Operator = "NOT"
	EQ    Operator = "EQ"
	NEQ   Operator = "NEQ"
	LT    Operator = "LT"
	LTE   Operator = "LTE"
	GT    Operator = "GT"
	GTE   Operator = "GTE"
	IN    Operator = "IN"
	RE    Operator = "RE"
)

type Query struct {
	Oper Operator `json:"operator,omitempty"`

	*QueryField
	*QueryMatch
	*QueryGroup
}

type Queries []*Query

type QueryField struct {
	Field string `json:"field"`
}

type QueryMatch struct {
	Value interface{} `json:"value,omitempty"`
}

type QueryGroup struct {
	Queries Queries `json:"queries"`
}

/*
 * Group
 */

func newGroup(q Queries) *Query {
	return &Query{
		Oper: GROUP,
		QueryGroup: &QueryGroup{
			Queries: q,
		},
	}
}

func Group(q Queries) Queries {
	return Queries{
		newGroup(q),
	}
}

func (q Queries) Group(sub Queries) Queries {
	q = append(q, newGroup(sub))
	return q
}

/*
 * Condition
 */

func newCondition(o Operator) *Query {
	return &Query{
		Oper: o,
	}
}

func (q Queries) And() Queries {
	q = append(q, newCondition(AND))
	return q
}

func (q Queries) Or() Queries {
	q = append(q, newCondition(OR))
	return q
}

func Not() Queries {
	return Queries{
		newCondition(NOT),
	}
}

func (q Queries) Not() Queries {
	i := len(q) - 1
	if q[i].QueryField != nil {
		q = append(q, q[i])
		q[i] = newCondition(NOT)
	} else {
		q[i] = newCondition(NOT)
	}
	return q
}

/*
 * Field
 */

func newField(k string) *Query {
	return &Query{
		QueryField: &QueryField{
			Field: k,
		},
	}
}

func Field(k string) Queries {
	return Queries{
		newField(k),
	}
}

func (q Queries) Field(k string) Queries {
	q = append(q, newField(k))
	return q
}

/*
 * Match
 */

func (q Queries) Last() *Query {
	return q[len(q)-1]
}

func (q Queries) Eq(v interface{}) Queries {
	last := q.Last()
	last.Oper = EQ
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) Neq(v interface{}) Queries {
	last := q.Last()
	last.Oper = NEQ
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) Lt(v interface{}) Queries {
	last := q.Last()
	last.Oper = LT
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) Lte(v interface{}) Queries {
	last := q.Last()
	last.Oper = LTE
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) Gt(v interface{}) Queries {
	last := q.Last()
	last.Oper = GT
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) Gte(v interface{}) Queries {
	last := q.Last()
	last.Oper = GTE
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) In(v interface{}) Queries {
	last := q.Last()
	last.Oper = IN
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

func (q Queries) Re(v interface{}) Queries {
	last := q.Last()
	last.Oper = RE
	last.QueryMatch = &QueryMatch{Value: v}
	return q
}

/*
 * Evaluate Query
 */

func (q Queries) Match(a interface{}) (bool, error) {
	//	for _, query := range q {
	//	}

	return false, nil
}
