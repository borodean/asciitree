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
	sorted := sortChildren(t, options)
	return sorted.Title + sorted.printChildren("")
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
	copy := *t
	copy.Children = append([]*Node(nil), copy.Children...)
	sort.SliceStable(copy.Children, func(i, j int) bool {
		a := copy.Children[i]
		b := copy.Children[j]
		if options.branchesFirst && a.isBranch() && !b.isBranch() {
			return true
		}
		if options.sortByTitle {
			return a.Title < b.Title
		}
		return false
	})
	for i, child := range copy.Children {
		copy.Children[i] = sortChildren(child, options)
	}
	return &copy
}
