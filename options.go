package asciitree

type SortOptions struct {
	dirsFirst bool
}

type SortOption interface {
	apply(*SortOptions)
}

func WithDirsFirst(value bool) SortOption {
	return dirsFirstOption(value)
}

type dirsFirstOption bool

func newSortOptions(opts ...SortOption) SortOptions {
	var options SortOptions
	for _, o := range opts {
		o.apply(&options)
	}
	return options
}

func (d dirsFirstOption) apply(opts *SortOptions) {
	opts.dirsFirst = bool(d)
}

// Verify that dirsFirstOption implements asciitree.SortOption
var _ SortOption = (*dirsFirstOption)(nil)
