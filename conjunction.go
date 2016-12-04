package gorules

//go:generate stringer -type=Conjunction

type Conjunction int

const (
	And Conjunction = iota
	Or
)
