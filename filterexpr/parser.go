package filterexpr

type NodeType int

const (
	NodeFilter NodeType = iota
	NodeOp
)

type Expr struct {
	Type   NodeType
	Op     string
	Filter RawFilter // (type, key, op, value)
	Left   *Expr
	Right  *Expr
}

type RawFilter struct {
	ValueType string
	Index     any
	Operator  string
	Value     string
}
