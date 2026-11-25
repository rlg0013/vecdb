package index

import "github.com/rlg0013/vecdb/internal/filter"

type Result struct {
	ID   uint64
	Dist float32
}

type SearchOpts struct {
	EfSearch int
	Filter   filter.Expr
}

type Index interface {
	Add(id uint64, v []float32, attrs map[string]any) error
	Delete(id uint64) error
	Search(q []float32, k int, opts SearchOpts) ([]Result, error)
	Len() int
}
