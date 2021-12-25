package asciitree

// SprintOption represents an option that can be provided to the Sprint method.
type SprintOption interface {
	apply(*sprintOptions)
}

type sprintOptions struct {
	branchesFirst bool
	sortByTitle   bool
}

type (
	branchesFirstOption bool
	sortByTitle         bool
)

// WithBranchesFirst is an option that makes the Sprint method print branches
// before leaves.
func WithBranchesFirst(value bool) SprintOption {
	return branchesFirstOption(value)
}

// WithSortByTitle is an option that makes the Sprint method print nodes in
// an alphanumerical order.
func WithSortByTitle(value bool) SprintOption {
	return sortByTitle(value)
}

func (d branchesFirstOption) apply(opts *sprintOptions) {
	opts.branchesFirst = bool(d)
}

func (s sortByTitle) apply(opts *sprintOptions) {
	opts.sortByTitle = bool(s)
}

func newSprintOptions(opts ...SprintOption) sprintOptions {
	var options sprintOptions
	for _, o := range opts {
		o.apply(&options)
	}
	return options
}

// Verify that branchesFirstOption implements asciitree.SprintOption
var _ SprintOption = (*branchesFirstOption)(nil)
