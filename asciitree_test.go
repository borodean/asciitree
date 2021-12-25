package asciitree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		give string
		want *Tree
	}{
		{"alfa", &Tree{title: "alfa"}},
		{"bravo", &Tree{title: "bravo"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := New(tt.give)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestNewBranch(t *testing.T) {
	tests := []struct {
		give string
		want *Tree
	}{
		{"alfa", &Tree{title: "alfa", forceBranch: true}},
		{"bravo", &Tree{title: "bravo", forceBranch: true}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := NewBranch(tt.give)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestSprint(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want string
	}{{
		name: "single node",
		tree: New("alfa"),
		want: `alfa`,
	}, {
		name: "branched nodes",
		tree: New("alfa").AddTrees(
			New("bravo.txt"),
			New("charlie").AddTrees(
				New("delta.txt"),
				New("echo").Add("foxtrot.txt", "golf.txt"),
			),
		),
		want: `alfa
├── bravo.txt
└── charlie
    ├── delta.txt
    └── echo
        ├── foxtrot.txt
        └── golf.txt`,
	}, {
		name: "intersected branched nodes",
		tree: New("alfa").AddTrees(
			New("bravo").Add("charlie.txt"),
			New("delta.txt"),
		),
		want: `alfa
├── bravo
│   └── charlie.txt
└── delta.txt`,
	}, {
		name: "multiline titles",
		tree: New("alfa\n[dir]\n[3 MB]").Add(
			"bravo.txt\n[file]\n[1 MB]",
			"charlie.txt\n[file]\n[2 MB]",
		),
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

func TestTreeAdd(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		give []string
		want *Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: []string{"bravo", "charlie"},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie"},
		}},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: []string{"charlie", "delta"},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie"},
			{title: "delta"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.Add(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAddBranches(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		give []string
		want *Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: []string{"bravo", "charlie"},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo", forceBranch: true},
			{title: "charlie", forceBranch: true},
		}},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: []string{"charlie", "delta"},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie", forceBranch: true},
			{title: "delta", forceBranch: true},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.AddBranches(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAddTrees(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		give []*Tree
		want *Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: []*Tree{New("bravo"), New("charlie")},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie"},
		}},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: []*Tree{New("charlie"), New("delta")},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie"},
			{title: "delta"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.AddTrees(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeIsBranch(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want bool
	}{
		{"empty", New("alfa"), false},
		{"has children", New("alfa").Add("bravo"), true},
		{"forced", NewBranch("alfa"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.IsBranch()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTreeNewChild(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Tree
		give      string
		wantTree  *Tree
		wantChild *Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: "bravo",
		wantTree: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
		}},
		wantChild: &Tree{title: "bravo"},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: "charlie",
		wantTree: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie"},
		}},
		wantChild: &Tree{title: "charlie"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChild := tt.tree.NewChild(tt.give)
			assert.DeepEqual(t, tt.tree, tt.wantTree, cmpOptions)
			assert.DeepEqual(t, gotChild, tt.wantChild, cmpOptions)
		})
	}
}

func TestTreeNewChildBranch(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Tree
		give      string
		wantTree  *Tree
		wantChild *Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: "bravo",
		wantTree: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo", forceBranch: true},
		}},
		wantChild: &Tree{title: "bravo", forceBranch: true},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: "charlie",
		wantTree: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo"},
			{title: "charlie", forceBranch: true},
		}},
		wantChild: &Tree{title: "charlie", forceBranch: true},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChild := tt.tree.NewChildBranch(tt.give)
			assert.DeepEqual(t, tt.tree, tt.wantTree, cmpOptions)
			assert.DeepEqual(t, gotChild, tt.wantChild, cmpOptions)
		})
	}
}

func TestTreeSetTitle(t *testing.T) {
	tests := []struct {
		tree *Tree
		give string
		want *Tree
	}{
		{New("alfa"), "bravo", &Tree{title: "bravo"}},
		{New("charlie"), "delta", &Tree{title: "delta"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := tt.tree.SetTitle(tt.give)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
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
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo", children: []*Tree{
				{title: "foxtrot", children: []*Tree{
					{title: "india.txt"},
				}},
				{title: "golf.txt"},
				{title: "hotel", forceBranch: true},
			}},
			{title: "charlie.txt"},
			{title: "delta", forceBranch: true},
			{title: "echo.txt"},
			{title: "kilo", children: []*Tree{
				{title: "juliet.txt"},
			}},
		}},
	}, {
		name: "directories before files",
		give: []SortOption{WithBranchesFirst(true)},
		want: &Tree{title: "alfa", children: []*Tree{
			{title: "bravo", children: []*Tree{
				{title: "foxtrot", children: []*Tree{
					{title: "india.txt"},
				}},
				{title: "hotel", forceBranch: true},
				{title: "golf.txt"},
			}},
			{title: "delta", forceBranch: true},
			{title: "kilo", children: []*Tree{
				{title: "juliet.txt"},
			}},
			{title: "charlie.txt"},
			{title: "echo.txt"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := New("alfa").AddTrees(
				New("charlie.txt"),
				New("bravo").AddTrees(
					New("golf.txt"),
					New("foxtrot").Add("india.txt"),
					NewBranch("hotel"),
				),
				New("kilo").Add("juliet.txt"),
				NewBranch("delta"),
				New("echo.txt"),
			)
			got := tree.Sort(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeTitle(t *testing.T) {
	tests := []struct {
		tree *Tree
		want string
	}{
		{New("alfa"), "alfa"},
		{New("bravo"), "bravo"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := tt.tree.Title()
			assert.Equal(t, got, tt.want)
		})
	}
}

var cmpOptions = cmp.AllowUnexported(Tree{})
