package asciitree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestTreeIsBranch(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want bool
	}{{
		name: "empty",
		tree: &Tree{Title: "alfa"},
		want: false,
	}, {
		name: "has children",
		tree: &Tree{Title: "alfa", Children: []*Tree{
			{Title: "bravo"},
		}},
		want: true,
	}, {
		name: "forced",
		tree: &Tree{Title: "alfa", ForceBranch: true},
		want: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.IsBranch()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTreeSort(t *testing.T) {
	tests := []struct {
		name string
		give []SortOption
		want *Tree
	}{{
		name: "default",
		give: []SortOption{},
		want: &Tree{Title: "alfa", Children: []*Tree{
			{Title: "bravo", Children: []*Tree{
				{Title: "foxtrot", Children: []*Tree{
					{Title: "india.txt"},
				}},
				{Title: "golf.txt"},
				{Title: "hotel", ForceBranch: true},
			}},
			{Title: "charlie.txt"},
			{Title: "delta", ForceBranch: true},
			{Title: "echo.txt"},
			{Title: "kilo", Children: []*Tree{
				{Title: "juliet.txt"},
			}},
		}},
	}, {
		name: "directories before files",
		give: []SortOption{WithBranchesFirst(true)},
		want: &Tree{Title: "alfa", Children: []*Tree{
			{Title: "bravo", Children: []*Tree{
				{Title: "foxtrot", Children: []*Tree{
					{Title: "india.txt"},
				}},
				{Title: "hotel", ForceBranch: true},
				{Title: "golf.txt"},
			}},
			{Title: "delta", ForceBranch: true},
			{Title: "kilo", Children: []*Tree{
				{Title: "juliet.txt"},
			}},
			{Title: "charlie.txt"},
			{Title: "echo.txt"},
		}},
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
			got := tree.Sort(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

var cmpOptions = cmp.AllowUnexported(Tree{})
