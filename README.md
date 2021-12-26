# asciitree

Package asciitree provides tools to build directory trees and print them
using ASCII art.

## Types

### type [Node](/asciitree.go#L12)

`type Node struct { ... }`

Node represents a directory tree node.

#### func [NewDir](/asciitree.go#L22)

`func NewDir(name string) *Node`

NewDir creates a tree node and forces it to be recognized as a directory.

#### func [NewFile](/asciitree.go#L27)

`func NewFile(name string) *Node`

NewFile creates a tree node.

#### func (*Node) [Add](/asciitree.go#L35)

`func (n *Node) Add(nodes ...*Node) *Node`

Add appends the provided tree nodes to the node's children.

Unlike AddFile and AddDir, Add returns the original node for
chaining.

#### func (*Node) [AddDir](/asciitree.go#L44)

`func (n *Node) AddDir(name string) *Node`

AddDir creates a tree node, forces it to be recognized as a directory,
and appends it to the node's children.

Unlike AddDirs, AddDir returns the newly created node.

#### func (*Node) [AddDirs](/asciitree.go#L54)

`func (n *Node) AddDirs(names ...string) *Node`

AddDirs creates one or more tree nodes with the provided names, forces them
to be recognized as directories, and appends them to the node's children.

Unlike AddDir, AddDirs returns the original node for chaining.

#### func (*Node) [AddFile](/asciitree.go#L66)

`func (n *Node) AddFile(name string) *Node`

AddFile creates a tree node and appends it to the node's children.

Unlike AddFiles, AddFile returns the newly created node.

#### func (*Node) [AddFiles](/asciitree.go#L76)

`func (n *Node) AddFiles(names ...string) *Node`

AddFiles creates one or more tree nodes with the provided names and appends
them to the node's children.

Unlike AddFile, AddFiles returns the original node for chaining.

#### func (*Node) [Sort](/asciitree.go#L88)

`func (n *Node) Sort(opts ...SortOption) *Node`

Sort recursively sorts the node's children in place.

Sort returns the original node for chaining.

#### func (*Node) [String](/asciitree.go#L94)

`func (n *Node) String() string`

String returns the tree's visual representation.

### type [SortOption](/options.go#L8)

`type SortOption interface { ... }`

SortOption represents an option that can be provided to the Sort method.

#### func [WithDirsFirst](/options.go#L20)

`func WithDirsFirst(value bool) SortOption`

WithDirsFirst is an option that makes the Sort method order directories
before leaves.

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
