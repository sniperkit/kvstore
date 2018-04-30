package qry

import (
	"reflect"

	"github.com/mickep76/kvstore/cmp"
)

type Operator string

const (
	EQ  Operator = "EQ"
	NEQ Operator = "NEQ"
	LT  Operator = "LT"
	LTE Operator = "LTE"
	GT  Operator = "GT"
	GTE Operator = "GTE"
	RE  Operator = "RE"
)

type Query struct {
	Operator Operator
	Field    string
	Value    interface{}
	Queries  Queries
}

type Queries []*Query

func NewQuery(operator Operator, field string, value interface{}) *Query {
	return &Query{
		Operator: operator,
		Field:    field,
		Value:    value,
	}
}

func Eq(field string, value interface{}) *Query {
	return &Query{
		Operator: EQ,
		Field:    field,
		Value:    value,
	}
}

func Neq(field string, value interface{}) *Query {
	return &Query{
		Operator: NEQ,
		Field:    field,
		Value:    value,
	}
}

func Lt(field string, value interface{}) *Query {
	return &Query{
		Operator: LT,
		Field:    field,
		Value:    value,
	}
}

func Lte(field string, value interface{}) *Query {
	return &Query{
		Operator: LTE,
		Field:    field,
		Value:    value,
	}
}

func Gt(field string, value interface{}) *Query {
	return &Query{
		Operator: GT,
		Field:    field,
		Value:    value,
	}
}

func Gte(field string, value interface{}) *Query {
	return &Query{
		Operator: GTE,
		Field:    field,
		Value:    value,
	}
}

func (q *Query) Match(a interface{}) ([]interface{}, error) {
	v := reflect.Indirect(reflect.ValueOf(a))
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return nil, ErrKindNotSupported
	}

	if v.Len() == 0 {
		return nil, nil
	}

	var results []interface{}
	for i := 0; i < v.Len(); i++ {
		fv, err := FieldValue(v.Index(i).Interface(), q.Field)
		if err != nil {
			return nil, err
		}

		var match bool
		switch q.Operator {
		case EQ:
			match, err = cmp.Eq(fv, q.Value)
		case NEQ:
			match, err = cmp.Neq(fv, q.Value)
		case LT:
			match, err = cmp.Lt(fv, q.Value)
		case LTE:
			match, err = cmp.Lte(fv, q.Value)
		case GT:
			match, err = cmp.Gt(fv, q.Value)
		case GTE:
			match, err = cmp.Gte(fv, q.Value)
		default:
			return nil, ErrUnknownOperator
		}

		if err != nil {
			return nil, err
		}
		if match {
			results = append(results, v.Index(i).Interface())
		}
	}

	return results, nil
}
