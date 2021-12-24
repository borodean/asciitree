package asciitree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		give string
		want Tree
	}{
		{"alfa", &tree{title: "alfa"}},
		{"bravo", &tree{title: "bravo"}},
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
		want Tree
	}{
		{"alfa", &tree{title: "alfa", forceBranch: true}},
		{"bravo", &tree{title: "bravo", forceBranch: true}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := NewBranch(tt.give)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAdd(t *testing.T) {
	tests := []struct {
		name string
		tree Tree
		give []string
		want Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: []string{"bravo", "charlie"},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie"},
		}},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: []string{"charlie", "delta"},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie"},
			&tree{title: "delta"},
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
		tree Tree
		give []string
		want Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: []string{"bravo", "charlie"},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo", forceBranch: true},
			&tree{title: "charlie", forceBranch: true},
		}},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: []string{"charlie", "delta"},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie", forceBranch: true},
			&tree{title: "delta", forceBranch: true},
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
		tree Tree
		give []Tree
		want Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: []Tree{New("bravo"), New("charlie")},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie"},
		}},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: []Tree{New("charlie"), New("delta")},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie"},
			&tree{title: "delta"},
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
		tree Tree
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
		tree      Tree
		give      string
		wantTree  Tree
		wantChild Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: "bravo",
		wantTree: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
		}},
		wantChild: &tree{title: "bravo"},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: "charlie",
		wantTree: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie"},
		}},
		wantChild: &tree{title: "charlie"},
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
		tree      Tree
		give      string
		wantTree  Tree
		wantChild Tree
	}{{
		name: "empty",
		tree: New("alfa"),
		give: "bravo",
		wantTree: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo", forceBranch: true},
		}},
		wantChild: &tree{title: "bravo", forceBranch: true},
	}, {
		name: "has children",
		tree: New("alfa").Add("bravo"),
		give: "charlie",
		wantTree: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo"},
			&tree{title: "charlie", forceBranch: true},
		}},
		wantChild: &tree{title: "charlie", forceBranch: true},
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
		tree Tree
		give string
		want Tree
	}{
		{New("alfa"), "bravo", &tree{title: "bravo"}},
		{New("charlie"), "delta", &tree{title: "delta"}},
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
		want Tree
	}{{
		name: "default",
		give: []SortOption{},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo", children: []Tree{
				&tree{title: "foxtrot", children: []Tree{
					&tree{title: "india.txt"},
				}},
				&tree{title: "golf.txt"},
				&tree{title: "hotel", forceBranch: true},
			}},
			&tree{title: "charlie.txt"},
			&tree{title: "delta", forceBranch: true},
			&tree{title: "echo.txt"},
			&tree{title: "kilo", children: []Tree{
				&tree{title: "juliet.txt"},
			}},
		}},
	}, {
		name: "directories before files",
		give: []SortOption{WithDirsFirst(true)},
		want: &tree{title: "alfa", children: []Tree{
			&tree{title: "bravo", children: []Tree{
				&tree{title: "foxtrot", children: []Tree{
					&tree{title: "india.txt"},
				}},
				&tree{title: "hotel", forceBranch: true},
				&tree{title: "golf.txt"},
			}},
			&tree{title: "delta", forceBranch: true},
			&tree{title: "kilo", children: []Tree{
				&tree{title: "juliet.txt"},
			}},
			&tree{title: "charlie.txt"},
			&tree{title: "echo.txt"},
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

func TestTreeString(t *testing.T) {
	tests := []struct {
		name string
		tree Tree
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
			got := tt.tree.String()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTreeTitle(t *testing.T) {
	tests := []struct {
		tree Tree
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

var cmpOptions = cmp.AllowUnexported(tree{})
