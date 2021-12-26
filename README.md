# asciitree

Package asciitree provides tools to build trees of entities and print them
using ASCII art.

## Types

### type [SortOption](/options.go#L4)

`type SortOption interface { ... }`

SortOption represents an option that can be provided to the Sort method.

#### func [WithBranchesFirst](/options.go#L16)

`func WithBranchesFirst(value bool) SortOption`

WithBranchesFirst is an option that makes the Sort method order branches
before leaves.

### type [Tree](/asciitree.go#L12)

`type Tree struct { ... }`

Tree represents a tree node.

#### func [New](/asciitree.go#L19)

`func New(title string) *Tree`

New creates a tree node.

#### func [NewBranch](/asciitree.go#L24)

`func NewBranch(title string) *Tree`

NewBranch creates a tree node and forces it to be recognized as a branch.

#### func (*Tree) [Add](/asciitree.go#L32)

`func (t *Tree) Add(titles ...string) *Tree`

Add creates one or more tree nodes with the provided titles and appends them
to the node's children.

Unlike NewChild, Add returns the original node for chaining.

#### func (*Tree) [AddBranches](/asciitree.go#L44)

`func (t *Tree) AddBranches(titles ...string) *Tree`

AddBranches creates one or more tree nodes with the provided titles, forces
them to be recognized as branches, and appends them to the node's children.

Unlike NewChildBranch, AddBranches returns the original node for chaining.

#### func (*Tree) [AddTrees](/asciitree.go#L56)

`func (t *Tree) AddTrees(trees ...*Tree) *Tree`

AddTrees appends the provided tree nodes to the node's children.

Unlike NewChild and NewChildBranch, AddTrees returns the original node for
chaining.

#### func (*Tree) [IsBranch](/asciitree.go#L69)

`func (t *Tree) IsBranch() bool`

IsBranch reports whether the node should be recognized as a branch. This is
possible in two cases:

- The node has one or more children.
- The node was forced to be recognized as a branch by creating it with the
NewBranch function or the NewChildBranch method.

#### func (*Tree) [NewChild](/asciitree.go#L76)

`func (t *Tree) NewChild(title string) *Tree`

NewChild creates a tree node and appends it to the node's children.

Unlike Add, NewChild returns the newly created node.

#### func (*Tree) [NewChildBranch](/asciitree.go#L86)

`func (t *Tree) NewChildBranch(title string) *Tree`

NewChildBranch creates a tree node, forces it to be recognized as a branch,
and appends it to the node's children.

Unlike AddBranches, NewChildBranch returns the newly created node.

#### func (*Tree) [SetTitle](/asciitree.go#L95)

`func (t *Tree) SetTitle(title string) *Tree`

SetTitle sets the node's title.

SetTitle returns the original node for chaining.

#### func (*Tree) [Sort](/asciitree.go#L103)

`func (t *Tree) Sort(opts ...SortOption) *Tree`

Sort recursively sorts the node's children in place.

Sort returns the original node for chaining.

#### func (*Tree) [String](/asciitree.go#L120)

`func (t *Tree) String() string`

String returns the tree's visual representation.

#### func (*Tree) [Title](/asciitree.go#L125)

`func (t *Tree) Title() string`

Title returns the node's title.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
