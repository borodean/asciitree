package asciitree

type SortOptions struct {
	branchesFirst bool
}

type SortOption interface {
	apply(*SortOptions)
}

func WithBranchesFirst(value bool) SortOption {
	return branchesFirstOption(value)
}

type branchesFirstOption bool

func newSortOptions(opts ...SortOption) SortOptions {
	var options SortOptions
	for _, o := range opts {
		o.apply(&options)
	}
	return options
}

func (d branchesFirstOption) apply(opts *SortOptions) {
	opts.branchesFirst = bool(d)
}

// Verify that branchesFirstOption implements asciitree.SortOption
var _ SortOption = (*branchesFirstOption)(nil)
