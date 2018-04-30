package qry

import (
	"reflect"

	"github.com/mickep76/kvstore/cmp"
)

// Support and, not, or
// Suppoty tag, orderBy, limit

type Operator int

const (
	AND Operator = iota
	OR
	NOT
	EQ
	NEQ
	LT
	LTE
	GT
	GTE
	RE
)

type Query struct {
	Tag     string
	OrderBy string
	Limit   int
	Matches Matches
}

type Match struct {
	Operator Operator
	Query    *Query
	Field    string
	Value    interface{}
}

type Matches []*Match

func NewQuery() *Query {
	return &Query{}
}

/*
func (q *Query) WithTag(tag string) *Query {
	q.tag = tag
	return q
}

func (q *Query) WithOrderBy(orderBy string) *Query {
	q.orderBy = orderBy
	return q
}

func (q *Query) WithLimit(limit int) *Query {
	q.limit = limit
	return q
}
*/

func (q *Query) AddMatchSubQuery(operator Operator, query *Query) *Query {
	q.Matches = append(q.Matches, &Match{
		Operator: operator,
		Query:    query,
	})
	return q
}

func (q *Query) AddMatch(operator Operator, field string, value interface{}) *Query {
	q.Matches = append(q.Matches, &Match{
		Operator: operator,
		Field:    field,
		Value:    value,
	})
	return q
}

func And(query *Query) *Query {
	return NewQuery().AddMatchSubQuery(AND, query)
}

func Or(query *Query) *Query {
	return NewQuery().AddMatchSubQuery(OR, query)
}

func Not(query *Query) *Query {
	return NewQuery().AddMatchSubQuery(NOT, query)
}

func Eq(field string, value interface{}) *Query {
	return NewQuery().AddMatch(EQ, field, value)
}

func Neq(field string, value interface{}) *Query {
	return NewQuery().AddMatch(NEQ, field, value)
}

func Lt(field string, value interface{}) *Query {
	return NewQuery().AddMatch(LT, field, value)
}

func Lte(field string, value interface{}) *Query {
	return NewQuery().AddMatch(LTE, field, value)
}

func Gt(field string, value interface{}) *Query {
	return NewQuery().AddMatch(GT, field, value)
}

func Gte(field string, value interface{}) *Query {
	return NewQuery().AddMatch(GTE, field, value)
}

func Re(field string, value string) *Query {
	return NewQuery().AddMatch(RE, field, value)
}

func (q *Query) And(query *Query) *Query {
	return q.AddMatchSubQuery(AND, query)
}

func (q *Query) Or(query *Query) *Query {
	return q.AddMatchSubQuery(OR, query)
}

func (q *Query) Not(query *Query) *Query {
	return q.AddMatchSubQuery(NOT, query)
}

func (q *Query) Eq(field string, value interface{}) *Query {
	return q.AddMatch(EQ, field, value)
}

func (q *Query) Neq(field string, value interface{}) *Query {
	return q.AddMatch(NEQ, field, value)
}

func (q *Query) Lt(field string, value interface{}) *Query {
	return q.AddMatch(LT, field, value)
}

func (q *Query) Lte(field string, value interface{}) *Query {
	return q.AddMatch(LTE, field, value)
}

func (q *Query) Gt(field string, value interface{}) *Query {
	return q.AddMatch(GT, field, value)
}

func (q *Query) Gte(field string, value interface{}) *Query {
	return q.AddMatch(GTE, field, value)
}

func (q *Query) Re(field string, value string) *Query {
	return q.AddMatch(GTE, field, value)
}

func (q *Query) Match(a interface{}) ([]interface{}, error) {
	first := true
	var r []interface{}
	for _, m := range q.Matches {
		var err error
		if first {
			r, err = m.Match(a)
			first = false
		} else {
			r, err = m.Match(r)
		}

		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

func (m *Match) Match(a interface{}) ([]interface{}, error) {
	v := reflect.Indirect(reflect.ValueOf(a))
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return nil, ErrKindNotSupported
	}

	var r []interface{}
	for i := 0; i < v.Len(); i++ {
		fv, err := FieldValue(v.Index(i).Interface(), m.Field)
		if err != nil {
			return nil, err
		}

		var matched bool
		switch m.Operator {
		case EQ:
			matched, err = cmp.Eq(fv, m.Value)
		case NEQ:
			matched, err = cmp.Neq(fv, m.Value)
		case LT:
			matched, err = cmp.Lt(fv, m.Value)
		case LTE:
			matched, err = cmp.Lte(fv, m.Value)
		case GT:
			matched, err = cmp.Gt(fv, m.Value)
		case GTE:
			matched, err = cmp.Gte(fv, m.Value)
		case RE:
			matched, err = cmp.Re(m.Value.(string), fv)
		}

		if err != nil {
			return nil, err
		}

		if matched {
			r = append(r, v.Index(i).Interface())
		}
	}

	return r, nil
}
