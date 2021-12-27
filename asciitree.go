// Package asciitree builds and prints directory trees using ASCII art.
package asciitree

import (
	"fmt"
	"sort"
	"strings"
)

// Node describes a directory tree node. It can be either a directory or a file,
// depending on the IsDir value.
type Node struct {
	// Name is the name of the directory or file described by the node.
	Name string
	// IsDir identifies whether the node describes a directory.
	IsDir bool
	// Children is a slice of all the immediate descendants of the node.
	Children []*Node
}

// NewDir creates a directory node.
func NewDir(name string) *Node {
	return &Node{Name: name, IsDir: true}
}

// NewFile creates a file node.
func NewFile(name string) *Node {
	return &Node{Name: name}
}

// Add places the given nodes under the current node and returns the current
// node.
func (n *Node) Add(nodes ...*Node) *Node {
	n.Children = append(n.Children, nodes...)
	return n
}

// AddDir creates a directory node under the current node and returns the newly
// created node.
func (n *Node) AddDir(name string) *Node {
	child := NewDir(name)
	n.Children = append(n.Children, child)
	return child
}

// AddDirs creates one or more directory nodes under the current node and
// returns the current node.
func (n *Node) AddDirs(names ...string) *Node {
	nodes := make([]*Node, len(names))
	for i, name := range names {
		nodes[i] = NewDir(name)
	}
	n.Add(nodes...)
	return n
}

// AddFile creates a file node under the current node and returns the newly
// created node.
func (n *Node) AddFile(name string) *Node {
	child := NewFile(name)
	n.Children = append(n.Children, child)
	return child
}

// AddFiles creates one or more file nodes under the current node and returns
// the current node.
func (n *Node) AddFiles(names ...string) *Node {
	nodes := make([]*Node, len(names))
	for i, name := range names {
		nodes[i] = NewFile(name)
	}
	n.Add(nodes...)
	return n
}

// Sort recursively sorts the node's descendants alphanumerically and returns
// the current node.
func (n *Node) Sort(opts ...SortOption) *Node {
	n.sort(newSortOptions(opts...))
	return n
}

// String returns the ASCII art representation of the directory tree described
// by the current node and its descendants.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString(n.Name)
	n.string(&builder, "")
	return builder.String()
}

func (n *Node) sort(options sortOptions) {
	if len(n.Children) == 0 {
		return
	}
	sort.SliceStable(n.Children, func(i, j int) bool {
		a := n.Children[i]
		b := n.Children[j]
		if options.dirsFirst && a.IsDir && !b.IsDir {
			return true
		}
		return a.Name < b.Name
	})
	for _, child := range n.Children {
		child.sort(options)
	}
}

func (n *Node) string(builder *strings.Builder, prefix string) {
	for i, child := range n.Children {
		connector := "├── "
		spacer := "│   "
		nspacer := "\n│   "
		if i == len(n.Children)-1 {
			connector = "└── "
			spacer = "    "
			nspacer = "\n    "
		}
		builder.WriteString("\n")
		builder.WriteString(prefix)
		builder.WriteString(connector)
		builder.WriteString(strings.ReplaceAll(child.Name, "\n", nspacer))
		child.string(builder, prefix+spacer)
	}
}

// Verify that Node implements fmt.Stringer:
var _ fmt.Stringer = (*Node)(nil)
