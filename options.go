package asciitree

// SortOption represents an option that can be provided to the Sort method.
type SortOption interface {
	apply(*sortOptions)
}

type sortOptions struct {
	dirsFirst bool
}

type dirsFirstOption bool

// WithDirsFirst is an option that makes the Sort method order directories
// before leaves.
func WithDirsFirst(value bool) SortOption {
	return dirsFirstOption(value)
}

func (d dirsFirstOption) apply(opts *sortOptions) {
	opts.dirsFirst = bool(d)
}

func newSortOptions(opts ...SortOption) sortOptions {
	var options sortOptions
	for _, o := range opts {
		o.apply(&options)
	}
	return options
}

// Verify that dirsFirstOption implements asciitree.SortOption
var _ SortOption = (*dirsFirstOption)(nil)
