package asciitree

// SortOption represents an option that can be provided to the Sort method.
type SortOption interface {
	apply(*sortOptions)
}

type sortOptions struct {
	branchesFirst bool
}

type branchesFirstOption bool

// WithBranchesFirst is an option that makes the Sort method order branches
// before leaves.
func WithBranchesFirst(value bool) SortOption {
	return branchesFirstOption(value)
}

func (d branchesFirstOption) apply(opts *sortOptions) {
	opts.branchesFirst = bool(d)
}

func newSortOptions(opts ...SortOption) sortOptions {
	var options sortOptions
	for _, o := range opts {
		o.apply(&options)
	}
	return options
}

// Verify that branchesFirstOption implements asciitree.SortOption
var _ SortOption = (*branchesFirstOption)(nil)
