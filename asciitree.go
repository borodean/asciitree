package asciitree

import (
	"fmt"
	"sort"
	"strings"
)

type Tree interface {
	fmt.Stringer

	Add(...string) Tree
	AddBranches(...string) Tree
	AddTrees(...Tree) Tree
	IsBranch() bool
	NewChild(string) Tree
	NewChildBranch(string) Tree
	SetTitle(string) Tree
	Sort(...SortOption) Tree
	Title() string

	printChildren(string) string
}

type tree struct {
	children    []Tree
	forceBranch bool
	title       string
}

func New(title string) Tree {
	return &tree{title: title}
}

func NewBranch(title string) Tree {
	return &tree{forceBranch: true, title: title}
}

func (t *tree) Add(titles ...string) Tree {
	for _, title := range titles {
		child := New(title)
		t.children = append(t.children, child)
	}
	return t
}

func (t *tree) AddBranches(titles ...string) Tree {
	for _, title := range titles {
		child := NewBranch(title)
		t.children = append(t.children, child)
	}
	return t
}

func (t *tree) AddTrees(trees ...Tree) Tree {
	for _, tree := range trees {
		t.children = append(t.children, tree)
	}
	return t
}

func (t *tree) IsBranch() bool {
	return t.forceBranch || len(t.children) > 0
}

func (t *tree) NewChild(title string) Tree {
	child := New(title)
	t.children = append(t.children, child)
	return child
}

func (t *tree) NewChildBranch(title string) Tree {
	child := NewBranch(title)
	t.children = append(t.children, child)
	return child
}

func (t *tree) SetTitle(title string) Tree {
	t.title = title
	return t
}

func (t *tree) Sort(opts ...SortOption) Tree {
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

func (t *tree) String() string {
	return t.Title() + t.printChildren("")
}

func (t *tree) Title() string {
	return t.title
}

func (t *tree) printChildren(prefix string) string {
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

// Verify that tree implements asciitree.Tree
var _ Tree = (*tree)(nil)
