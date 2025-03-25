package filter

import "time"

type Operator string

const (
	OperatorContain            Operator = "CONTAIN"
	OperatorEqual              Operator = "EQUAL"
	OperatorNotEqual           Operator = "NOT_EQUAL"
	OperatorLessThan           Operator = "LESS_THAN"
	OperatorGreaterThan        Operator = "GREATER_THAN"
	OperatorLessThanOrEqual    Operator = "LESS_THAN_OR_EQUAL"
	OperatorGreaterThanOrEqual Operator = "GREATER_THAN_OR_EQUAL"
)

type Condition string

const (
	ConditionAnd Condition = "AND"
	ConditionOr  Condition = "OR"
)

type ValueType string

const (
	ValueTypeNumber   ValueType = "NUMBER"
	ValueTypeString   ValueType = "STRING"
	ValueTypeDatetime ValueType = "DATETIME"
)

type ValueNumber interface {
	int | float64 | float32
}

type Value interface {
	ValueNumber | string | time.Time
}
