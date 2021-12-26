package asciitree

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestNewDir(t *testing.T) {
	tests := []struct {
		give string
		want *Tree
	}{
		{"alfa", &Tree{name: "alfa", forceDir: true}},
		{"bravo", &Tree{name: "bravo", forceDir: true}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := NewDir(tt.give)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestNewFile(t *testing.T) {
	tests := []struct {
		give string
		want *Tree
	}{
		{"alfa", &Tree{name: "alfa"}},
		{"bravo", &Tree{name: "bravo"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := NewFile(tt.give)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAdd(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		give []*Tree
		want *Tree
	}{{
		name: "empty",
		tree: NewFile("alfa"),
		give: []*Tree{NewFile("bravo"), NewFile("charlie")},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie"},
		}},
	}, {
		name: "has children",
		tree: NewFile("alfa").AddFiles("bravo"),
		give: []*Tree{NewFile("charlie"), NewFile("delta")},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie"},
			{name: "delta"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.Add(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAddDirs(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		give []string
		want *Tree
	}{{
		name: "empty",
		tree: NewFile("alfa"),
		give: []string{"bravo", "charlie"},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo", forceDir: true},
			{name: "charlie", forceDir: true},
		}},
	}, {
		name: "has children",
		tree: NewFile("alfa").AddFiles("bravo"),
		give: []string{"charlie", "delta"},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie", forceDir: true},
			{name: "delta", forceDir: true},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.AddDirs(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAddFiles(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		give []string
		want *Tree
	}{{
		name: "empty",
		tree: NewFile("alfa"),
		give: []string{"bravo", "charlie"},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie"},
		}},
	}, {
		name: "has children",
		tree: NewFile("alfa").AddFiles("bravo"),
		give: []string{"charlie", "delta"},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie"},
			{name: "delta"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.AddFiles(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeIsDir(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want bool
	}{
		{"empty", NewFile("alfa"), false},
		{"has children", NewFile("alfa").AddFiles("bravo"), true},
		{"forced", NewDir("alfa"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.IsDir()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTreeName(t *testing.T) {
	tests := []struct {
		tree *Tree
		want string
	}{
		{NewFile("alfa"), "alfa"},
		{NewFile("bravo"), "bravo"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := tt.tree.Name()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTreeNewChildDir(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Tree
		give      string
		wantTree  *Tree
		wantChild *Tree
	}{{
		name: "empty",
		tree: NewFile("alfa"),
		give: "bravo",
		wantTree: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo", forceDir: true},
		}},
		wantChild: &Tree{name: "bravo", forceDir: true},
	}, {
		name: "has children",
		tree: NewFile("alfa").AddFiles("bravo"),
		give: "charlie",
		wantTree: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie", forceDir: true},
		}},
		wantChild: &Tree{name: "charlie", forceDir: true},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChild := tt.tree.NewChildDir(tt.give)
			assert.DeepEqual(t, tt.tree, tt.wantTree, cmpOptions)
			assert.DeepEqual(t, gotChild, tt.wantChild, cmpOptions)
		})
	}
}

func TestTreeNewChildFile(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Tree
		give      string
		wantTree  *Tree
		wantChild *Tree
	}{{
		name: "empty",
		tree: NewFile("alfa"),
		give: "bravo",
		wantTree: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
		}},
		wantChild: &Tree{name: "bravo"},
	}, {
		name: "has children",
		tree: NewFile("alfa").AddFiles("bravo"),
		give: "charlie",
		wantTree: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo"},
			{name: "charlie"},
		}},
		wantChild: &Tree{name: "charlie"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChild := tt.tree.NewChildFile(tt.give)
			assert.DeepEqual(t, tt.tree, tt.wantTree, cmpOptions)
			assert.DeepEqual(t, gotChild, tt.wantChild, cmpOptions)
		})
	}
}

func TestTreeSetName(t *testing.T) {
	tests := []struct {
		tree *Tree
		give string
		want *Tree
	}{
		{NewFile("alfa"), "bravo", &Tree{name: "bravo"}},
		{NewFile("charlie"), "delta", &Tree{name: "delta"}},
	}
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			got := tt.tree.SetName(tt.give)
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
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo", children: []*Tree{
				{name: "foxtrot", children: []*Tree{
					{name: "india.txt"},
				}},
				{name: "golf.txt"},
				{name: "hotel", forceDir: true},
			}},
			{name: "charlie.txt"},
			{name: "delta", forceDir: true},
			{name: "echo.txt"},
			{name: "kilo", children: []*Tree{
				{name: "juliet.txt"},
			}},
		}},
	}, {
		name: "directories before files",
		give: []SortOption{WithDirsFirst(true)},
		want: &Tree{name: "alfa", children: []*Tree{
			{name: "bravo", children: []*Tree{
				{name: "foxtrot", children: []*Tree{
					{name: "india.txt"},
				}},
				{name: "hotel", forceDir: true},
				{name: "golf.txt"},
			}},
			{name: "delta", forceDir: true},
			{name: "kilo", children: []*Tree{
				{name: "juliet.txt"},
			}},
			{name: "charlie.txt"},
			{name: "echo.txt"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewFile("alfa").Add(
				NewFile("charlie.txt"),
				NewFile("bravo").Add(
					NewFile("golf.txt"),
					NewFile("foxtrot").AddFiles("india.txt"),
					NewDir("hotel"),
				),
				NewFile("kilo").AddFiles("juliet.txt"),
				NewDir("delta"),
				NewFile("echo.txt"),
			)
			got := tree.Sort(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeString(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want string
	}{{
		name: "just root",
		tree: NewFile("alfa"),
		want: `alfa`,
	}, {
		name: "nested",
		tree: NewFile("alfa").Add(
			NewFile("bravo.txt"),
			NewFile("charlie").Add(
				NewFile("delta.txt"),
				NewFile("echo").AddFiles("foxtrot.txt", "golf.txt"),
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
		name: "nested + intersected",
		tree: NewFile("alfa").Add(
			NewFile("bravo").AddFiles("charlie.txt"),
			NewFile("delta.txt"),
		),
		want: `alfa
├── bravo
│   └── charlie.txt
└── delta.txt`,
	}, {
		name: "multiline names",
		tree: NewFile("alfa\n[dir]\n[3 MB]").AddFiles(
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

var cmpOptions = cmp.AllowUnexported(Tree{})
