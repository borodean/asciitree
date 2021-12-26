// Package asciitree provides tools to build directory trees and print them
// using ASCII art.
package asciitree

import (
	"fmt"
	"sort"
	"strings"
)

// Node represents a directory tree node.
type Node struct {
	// Name is the name of the file (or directory) described by the node.
	Name string
	// IsDir identifies whether the node describes a directory.
	IsDir bool
	// Children is the slice of the node's children.
	Children []*Node
}

// NewDir creates a tree node and forces it to be recognized as a directory.
func NewDir(name string) *Node {
	return &Node{Name: name, IsDir: true}
}

// NewFile creates a tree node.
func NewFile(name string) *Node {
	return &Node{Name: name}
}

// Add appends the provided tree nodes to the node's children.
//
// Unlike AddFile and AddDir, Add returns the original node for
// chaining.
func (n *Node) Add(nodes ...*Node) *Node {
	n.Children = append(n.Children, nodes...)
	return n
}

// AddDir creates a tree node, forces it to be recognized as a directory,
// and appends it to the node's children.
//
// Unlike AddDirs, AddDir returns the newly created node.
func (n *Node) AddDir(name string) *Node {
	child := NewDir(name)
	n.Children = append(n.Children, child)
	return child
}

// AddDirs creates one or more tree nodes with the provided names, forces them
// to be recognized as directories, and appends them to the node's children.
//
// Unlike AddDir, AddDirs returns the original node for chaining.
func (n *Node) AddDirs(names ...string) *Node {
	nodes := make([]*Node, len(names))
	for i, name := range names {
		nodes[i] = NewDir(name)
	}
	n.Add(nodes...)
	return n
}

// AddFile creates a tree node and appends it to the node's children.
//
// Unlike AddFiles, AddFile returns the newly created node.
func (n *Node) AddFile(name string) *Node {
	child := NewFile(name)
	n.Children = append(n.Children, child)
	return child
}

// AddFiles creates one or more tree nodes with the provided names and appends
// them to the node's children.
//
// Unlike AddFile, AddFiles returns the original node for chaining.
func (n *Node) AddFiles(names ...string) *Node {
	nodes := make([]*Node, len(names))
	for i, name := range names {
		nodes[i] = NewFile(name)
	}
	n.Add(nodes...)
	return n
}

// Sort recursively sorts the node's children in place.
//
// Sort returns the original node for chaining.
func (n *Node) Sort(opts ...SortOption) *Node {
	n.sort(newSortOptions(opts...))
	return n
}

// String returns the tree's visual representation.
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

// Verify that Tree implements fmt.Stringer:
var _ fmt.Stringer = (*Node)(nil)
