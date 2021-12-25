// Package asciitree provides tools to build trees of entities and print them
// using ASCII art.
package asciitree

import (
	"sort"
	"strings"
)

// Node represents a tree node.
type Node struct {
	// Children is the slice of the node's children.
	Children []*Node
	// ForceBranch reports whether the node should be forced to be recognized as a
	// branch.
	ForceBranch bool
	// Title is the title of the node.
	Title string
}

// Sprint returns the node's visual representation.
func Sprint(t *Node, opts ...SprintOption) string {
	options := newSprintOptions(opts...)
	sortChildren(t, options)
	return t.Title + t.printChildren("")
}

func (t *Node) isBranch() bool {
	return t.ForceBranch || len(t.Children) > 0
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
			strings.ReplaceAll(child.Title, "\n", "\n"+spacer) +
			child.printChildren(prefix+spacer)
	}
	return out
}

func sortChildren(t *Node, options sprintOptions) *Node {
	sort.SliceStable(t.Children, func(i, j int) bool {
		a := t.Children[i]
		b := t.Children[j]
		if options.branchesFirst && a.isBranch() && !b.isBranch() {
			return true
		}
		if options.sortByTitle {
			return a.Title < b.Title
		}
		return false
	})
	for _, child := range t.Children {
		sortChildren(child, options)
	}
	return t
}
