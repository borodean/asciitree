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
	isDir    bool
	children []*Tree
}

// NewDir creates a tree node and forces it to be recognized as a directory.
func NewDir(name string) *Tree {
	return &Tree{isDir: true, name: name}
}

// NewFile creates a tree node.
func NewFile(name string) *Tree {
	return &Tree{name: name}
}

// Add appends the provided tree nodes to the node's children.
//
// Unlike AddFile and AddDir, Add returns the original node for
// chaining.
func (t *Tree) Add(trees ...*Tree) *Tree {
	for _, tree := range trees {
		t.children = append(t.children, tree)
	}
	return t
}

// AddDir creates a tree node, forces it to be recognized as a directory,
// and appends it to the node's children.
//
// Unlike AddDirs, AddDir returns the newly created node.
func (t *Tree) AddDir(name string) *Tree {
	child := NewDir(name)
	t.children = append(t.children, child)
	return child
}

// AddDirs creates one or more tree nodes with the provided names, forces them
// to be recognized as directories, and appends them to the node's children.
//
// Unlike AddDir, AddDirs returns the original node for chaining.
func (t *Tree) AddDirs(names ...string) *Tree {
	for _, name := range names {
		child := NewDir(name)
		t.children = append(t.children, child)
	}
	return t
}

// AddFile creates a tree node and appends it to the node's children.
//
// Unlike AddFiles, AddFile returns the newly created node.
func (t *Tree) AddFile(name string) *Tree {
	child := NewFile(name)
	t.children = append(t.children, child)
	return child
}

// AddFiles creates one or more tree nodes with the provided names and appends
// them to the node's children.
//
// Unlike AddFile, AddFiles returns the original node for chaining.
func (t *Tree) AddFiles(names ...string) *Tree {
	for _, name := range names {
		child := NewFile(name)
		t.children = append(t.children, child)
	}
	return t
}

// Name returns the node's name.
func (t *Tree) Name() string {
	return t.name
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
		if options.dirsFirst && a.isDir && !b.isDir {
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
