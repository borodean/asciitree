package asciitree

import (
	"fmt"
	"sort"
	"strings"
)

// Tree represents a tree node.
type Tree struct {
	children    []*Tree
	forceBranch bool
	title       string
}

// New creates a tree node.
func New(title string) *Tree {
	return &Tree{title: title}
}

// NewBranch creates a tree node and forces it to be recognized as a branch.
func NewBranch(title string) *Tree {
	return &Tree{forceBranch: true, title: title}
}

// Add creates one or more tree nodes with the provided titles and appends them
// to the node's children.
//
// Unlike NewChild, Add returns the original node for chaining.
func (t *Tree) Add(titles ...string) *Tree {
	for _, title := range titles {
		child := New(title)
		t.children = append(t.children, child)
	}
	return t
}

// AddBranches creates one or more tree nodes with the provided titles, forces
// them to be recognized as branches, and appends them to the node's children.
//
// Unlike NewChildBranch, AddBranches returns the original node for chaining.
func (t *Tree) AddBranches(titles ...string) *Tree {
	for _, title := range titles {
		child := NewBranch(title)
		t.children = append(t.children, child)
	}
	return t
}

// AddTrees appends the provided tree nodes to the node's children.
//
// Unlike NewChild and NewChildBranch, AddTrees returns the original node for
// chaining.
func (t *Tree) AddTrees(trees ...*Tree) *Tree {
	for _, tree := range trees {
		t.children = append(t.children, tree)
	}
	return t
}

// IsBranch reports whether the node should be recognized as a branch. This is
// possible in two cases:
//
// - The node has one or more children.
// - The node was forced to be recognized as a branch by creating it with the
// NewBranch function or the NewChildBranch method.
func (t *Tree) IsBranch() bool {
	return t.forceBranch || len(t.children) > 0
}

// NewChild creates a tree node and appends it to the node's children.
//
// Unlike Add, NewChild returns the newly created node.
func (t *Tree) NewChild(title string) *Tree {
	child := New(title)
	t.children = append(t.children, child)
	return child
}

// NewChildBranch creates a tree node, forces it to be recognized as a branch,
// and appends it to the node's children.
//
// Unlike AddBranches, NewChildBranch returns the newly created node.
func (t *Tree) NewChildBranch(title string) *Tree {
	child := NewBranch(title)
	t.children = append(t.children, child)
	return child
}

// SetTitle sets the node's title.
//
// SetTitle returns the original node for chaining.
func (t *Tree) SetTitle(title string) *Tree {
	t.title = title
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
		if options.dirsFirst && a.IsBranch() && !b.IsBranch() {
			return true
		}
		return a.Title() < b.Title()
	})
	for _, child := range t.children {
		child.Sort(opts...)
	}
	return t
}

// String returns the tree's visual representation.
func (t *Tree) String() string {
	return t.Title() + t.printChildren("")
}

// Title returns the node's title.
func (t *Tree) Title() string {
	return t.title
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
			strings.ReplaceAll(child.Title(), "\n", "\n"+spacer) +
			child.printChildren(prefix+spacer)
	}
	return out
}

// Verify that Tree implements fmt.Stringer:
var _ fmt.Stringer = (*Tree)(nil)
