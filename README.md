# asciitree

Package asciitree provides tools to build trees of entities and print them
using ASCII art.

## Functions

### func [Sprint](/asciitree.go#L22)

`func Sprint(node *Node, opts ...SprintOption) string`

Sprint returns the node's visual representation.

## Types

### type [Node](/asciitree.go#L11)

`type Node struct { ... }`

Node represents a tree node.

### type [SprintOption](/options.go#L4)

`type SprintOption interface { ... }`

SprintOption represents an option that can be provided to the Sprint method.

#### func [WithBranchesFirst](/options.go#L20)

`func WithBranchesFirst(value bool) SprintOption`

WithBranchesFirst is an option that makes the Sprint method print branches
before leaves.

#### func [WithSortByTitle](/options.go#L26)

`func WithSortByTitle(value bool) SprintOption`

WithSortByTitle is an option that makes the Sprint method print nodes in
an alphanumerical order.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
