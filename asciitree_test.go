package asciitree

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSprint(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want string
	}{{
		name: "single node",
		tree: &Tree{Title: "alfa"},
		want: `alfa`,
	}, {
		name: "branched nodes",
		tree: &Tree{Title: "alfa", Children: []*Tree{
			{Title: "bravo.txt"},
			{Title: "charlie", Children: []*Tree{
				{Title: "delta.txt"},
				{Title: "echo", Children: []*Tree{
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
		tree: &Tree{Title: "alfa", Children: []*Tree{
			{Title: "bravo", Children: []*Tree{
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
		tree: &Tree{Title: "alfa\n[dir]\n[3 MB]", Children: []*Tree{
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
		name: "default",
		give: []SprintOption{},
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
		name: "directories before files",
		give: []SprintOption{WithBranchesFirst(true)},
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
			tree := &Tree{Title: "alfa", Children: []*Tree{
				{Title: "charlie.txt"},
				{Title: "bravo", Children: []*Tree{
					{Title: "golf.txt"},
					{Title: "foxtrot", Children: []*Tree{
						{Title: "india.txt"},
					}},
					{Title: "hotel", ForceBranch: true},
				}},
				{Title: "kilo", Children: []*Tree{
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
