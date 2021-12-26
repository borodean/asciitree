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
func (t *Node) Add(nodes ...*Node) *Node {
	for _, tree := range nodes {
		t.Children = append(t.Children, tree)
	}
	return t
}

// AddDir creates a tree node, forces it to be recognized as a directory,
// and appends it to the node's children.
//
// Unlike AddDirs, AddDir returns the newly created node.
func (t *Node) AddDir(name string) *Node {
	child := NewDir(name)
	t.Children = append(t.Children, child)
	return child
}

// AddDirs creates one or more tree nodes with the provided names, forces them
// to be recognized as directories, and appends them to the node's children.
//
// Unlike AddDir, AddDirs returns the original node for chaining.
func (t *Node) AddDirs(names ...string) *Node {
	for _, name := range names {
		child := NewDir(name)
		t.Children = append(t.Children, child)
	}
	return t
}

// AddFile creates a tree node and appends it to the node's children.
//
// Unlike AddFiles, AddFile returns the newly created node.
func (t *Node) AddFile(name string) *Node {
	child := NewFile(name)
	t.Children = append(t.Children, child)
	return child
}

// AddFiles creates one or more tree nodes with the provided names and appends
// them to the node's children.
//
// Unlike AddFile, AddFiles returns the original node for chaining.
func (t *Node) AddFiles(names ...string) *Node {
	for _, name := range names {
		child := NewFile(name)
		t.Children = append(t.Children, child)
	}
	return t
}

// Sort recursively sorts the node's children in place.
//
// Sort returns the original node for chaining.
func (t *Node) Sort(opts ...SortOption) *Node {
	options := newSortOptions(opts...)
	sort.SliceStable(t.Children, func(i, j int) bool {
		a := t.Children[i]
		b := t.Children[j]
		if options.dirsFirst && a.IsDir && !b.IsDir {
			return true
		}
		return a.Name < b.Name
	})
	for _, child := range t.Children {
		child.Sort(opts...)
	}
	return t
}

// String returns the tree's visual representation.
func (t *Node) String() string {
	return t.Name + t.printChildren("")
}

func (t *Node) printChildren(prefix string) string {
	var out string
	for i, child := range t.Children {
		connector := "├── "
		spacer := "│   "
		if i == len(t.Children)-1 {
			connector = "└── "
			spacer = "    "
		}
		out += "\n" +
			prefix +
			connector +
			strings.ReplaceAll(child.Name, "\n", "\n"+spacer) +
			child.printChildren(prefix+spacer)
	}
	return out
}

// Verify that Tree implements fmt.Stringer:
var _ fmt.Stringer = (*Node)(nil)
