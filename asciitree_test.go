package asciitree

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSprint(t *testing.T) {
	tests := []struct {
		name string
		tree *Node
		want string
	}{{
		name: "single node",
		tree: &Node{Title: "alfa"},
		want: `alfa`,
	}, {
		name: "branched nodes",
		tree: &Node{Title: "alfa", Children: []*Node{
			{Title: "bravo.txt"},
			{Title: "charlie", Children: []*Node{
				{Title: "delta.txt"},
				{Title: "echo", Children: []*Node{
					{Title: "foxtrot.txt"},
					{Title: "golf.txt"},
				}},
			}},
		}},
		want: `alfa
├── bravo.txt
└── charlie
    ├── delta.txt
    └── echo
        ├── foxtrot.txt
        └── golf.txt`,
	}, {
		name: "intersected branched nodes",
		tree: &Node{Title: "alfa", Children: []*Node{
			{Title: "bravo", Children: []*Node{
				{Title: "charlie.txt"},
			}},
			{Title: "delta.txt"},
		}},
		want: `alfa
├── bravo
│   └── charlie.txt
└── delta.txt`,
	}, {
		name: "multiline titles",
		tree: &Node{Title: "alfa\n[dir]\n[3 MB]", Children: []*Node{
			{Title: "bravo.txt\n[file]\n[1 MB]"},
			{Title: "charlie.txt\n[file]\n[2 MB]"},
		}},
		want: `alfa
[dir]
[3 MB]
├── bravo.txt
│   [file]
│   [1 MB]
└── charlie.txt
    [file]
    [2 MB]`,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sprint(tt.tree)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestSprint_options(t *testing.T) {
	tests := []struct {
		name string
		give []SprintOption
		want string
	}{{
		name: "none",
		give: []SprintOption{},
		want: `alfa
├── charlie.txt
├── bravo
│   ├── golf.txt
│   ├── foxtrot
│   │   └── india.txt
│   └── hotel
├── kilo
│   └── juliet.txt
├── delta
└── echo.txt`,
	}, {
		name: "branches first",
		give: []SprintOption{WithBranchesFirst(true)},
		want: `alfa
├── bravo
│   ├── foxtrot
│   │   └── india.txt
│   ├── hotel
│   └── golf.txt
├── kilo
│   └── juliet.txt
├── delta
├── charlie.txt
└── echo.txt`,
	}, {
		name: "sort by title",
		give: []SprintOption{WithSortByTitle(true)},
		want: `alfa
├── bravo
│   ├── foxtrot
│   │   └── india.txt
│   ├── golf.txt
│   └── hotel
├── charlie.txt
├── delta
├── echo.txt
└── kilo
    └── juliet.txt`,
	}, {
		name: "branches first + sort by title",
		give: []SprintOption{WithBranchesFirst(true), WithSortByTitle(true)},
		want: `alfa
├── bravo
│   ├── foxtrot
│   │   └── india.txt
│   ├── hotel
│   └── golf.txt
├── delta
├── kilo
│   └── juliet.txt
├── charlie.txt
└── echo.txt`,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &Node{Title: "alfa", Children: []*Node{
				{Title: "charlie.txt"},
				{Title: "bravo", Children: []*Node{
					{Title: "golf.txt"},
					{Title: "foxtrot", Children: []*Node{
						{Title: "india.txt"},
					}},
					{Title: "hotel", ForceBranch: true},
				}},
				{Title: "kilo", Children: []*Node{
					{Title: "juliet.txt"},
				}},
				{Title: "delta", ForceBranch: true},
				{Title: "echo.txt"},
			}}
			got := Sprint(tree, tt.give...)
			assert.Equal(t, got, tt.want)
		})
	}
}
