package filter

type Expr interface {
	Eval(id uint64) bool
}
