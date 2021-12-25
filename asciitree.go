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
func Sprint(node *Node, opts ...SprintOption) string {
	options := newSprintOptions(opts...)
	sorted := sortChildren(node, options)
	return sorted.Title + printChildren(sorted, "")
}

func isBranch(node *Node) bool {
	return node.ForceBranch || len(node.Children) > 0
}

func printChildren(node *Node, prefix string) string {
	var out string
	for i, child := range node.Children {
		connector := "├── "
		spacer := "│   "
		if i == len(node.Children)-1 {
			connector = "└── "
			spacer = "    "
		}
		out += "\n" +
			prefix +
			connector +
			strings.ReplaceAll(child.Title, "\n", "\n"+spacer) +
			printChildren(child, prefix+spacer)
	}
	return out
}

func sortChildren(node *Node, options sprintOptions) *Node {
	copy := *node
	copy.Children = append([]*Node(nil), copy.Children...)
	sort.SliceStable(copy.Children, func(i, j int) bool {
		a := copy.Children[i]
		b := copy.Children[j]
		if options.branchesFirst && isBranch(a) && !isBranch(b) {
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
