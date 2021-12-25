package asciitree

// SprintOption represents an option that can be provided to the Sprint method.
type SprintOption interface {
	apply(*sprintOptions)
}

type sprintOptions struct {
	branchesFirst bool
}

type branchesFirstOption bool

// WithBranchesFirst is an option that makes the Sprint method print branches
// before leaves.
func WithBranchesFirst(value bool) SprintOption {
	return branchesFirstOption(value)
}

func (d branchesFirstOption) apply(opts *sprintOptions) {
	opts.branchesFirst = bool(d)
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
