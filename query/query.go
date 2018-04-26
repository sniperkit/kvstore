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
