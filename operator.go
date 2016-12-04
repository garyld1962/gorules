package gorules

//go:generate stringer -type=Operator
type Operator int

const (
	IsEqualTo  Operator    = iota
	IsNotEqualTo
	IsGreaterThan
	In
	NotIn
	IsLessThan
	IsGreaterThanOrEqualTo
	IsLessThanOrEqualTo
	IsLike
	IsNotLike
	IsNull
	IsNotNull
)
