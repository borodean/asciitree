// Package asciitree provides tools to build trees of entities and print them
// using ASCII art.
package asciitree

import (
	"fmt"
	"sort"
	"strings"
)

// Tree represents a tree node.
type Tree struct {
	name     string
	forceDir bool
	children []*Tree
}

// New creates a tree node.
func New(name string) *Tree {
	return &Tree{name: name}
}

// NewDir creates a tree node and forces it to be recognized as a directory.
func NewDir(name string) *Tree {
	return &Tree{forceDir: true, name: name}
}

// Add creates one or more tree nodes with the provided names and appends them
// to the node's children.
//
// Unlike NewChild, Add returns the original node for chaining.
func (t *Tree) Add(names ...string) *Tree {
	for _, name := range names {
		child := New(name)
		t.children = append(t.children, child)
	}
	return t
}

// AddDirs creates one or more tree nodes with the provided names, forces them
// to be recognized as directories, and appends them to the node's children.
//
// Unlike NewChildDir, AddDirs returns the original node for chaining.
func (t *Tree) AddDirs(names ...string) *Tree {
	for _, name := range names {
		child := NewDir(name)
		t.children = append(t.children, child)
	}
	return t
}

// AddTrees appends the provided tree nodes to the node's children.
//
// Unlike NewChild and NewChildDir, AddTrees returns the original node for
// chaining.
func (t *Tree) AddTrees(trees ...*Tree) *Tree {
	for _, tree := range trees {
		t.children = append(t.children, tree)
	}
	return t
}

// IsDir reports whether the node should be recognized as a directory. This is
// possible in two cases:
//
// - The node has one or more children.
// - The node was forced to be recognized as a directory by creating it with the
// NewDir function or the NewChildDir method.
func (t *Tree) IsDir() bool {
	return t.forceDir || len(t.children) > 0
}

// NewChild creates a tree node and appends it to the node's children.
//
// Unlike Add, NewChild returns the newly created node.
func (t *Tree) NewChild(name string) *Tree {
	child := New(name)
	t.children = append(t.children, child)
	return child
}

// NewChildDir creates a tree node, forces it to be recognized as a directory,
// and appends it to the node's children.
//
// Unlike AddDirs, NewChildDir returns the newly created node.
func (t *Tree) NewChildDir(name string) *Tree {
	child := NewDir(name)
	t.children = append(t.children, child)
	return child
}

// SetName sets the node's name.
//
// SetName returns the original node for chaining.
func (t *Tree) SetName(name string) *Tree {
	t.name = name
	return t
}

// Sort recursively sorts the node's children in place.
//
// Sort returns the original node for chaining.
func (t *Tree) Sort(opts ...SortOption) *Tree {
	options := newSortOptions(opts...)
	sort.SliceStable(t.children, func(i, j int) bool {
		a := t.children[i]
		b := t.children[j]
		if options.dirsFirst && a.IsDir() && !b.IsDir() {
			return true
		}
		return a.Name() < b.Name()
	})
	for _, child := range t.children {
		child.Sort(opts...)
	}
	return t
}

// String returns the tree's visual representation.
func (t *Tree) String() string {
	return t.Name() + t.printChildren("")
}

// Name returns the node's name.
func (t *Tree) Name() string {
	return t.name
}

func (t *Tree) printChildren(prefix string) string {
	var out string
	for i, child := range t.children {
		connector := "├── "
		spacer := "│   "
		if i == len(t.children)-1 {
			connector = "└── "
			spacer = "    "
		}
		out += "\n" +
			prefix +
			connector +
			strings.ReplaceAll(child.Name(), "\n", "\n"+spacer) +
			child.printChildren(prefix+spacer)
	}
	return out
}

// Verify that Tree implements fmt.Stringer:
var _ fmt.Stringer = (*Tree)(nil)
