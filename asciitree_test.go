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
		{"alfa", &Tree{Name: "alfa", IsDir: true}},
		{"bravo", &Tree{Name: "bravo", IsDir: true}},
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
		{"alfa", &Tree{Name: "alfa"}},
		{"bravo", &Tree{Name: "bravo"}},
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
		tree: NewDir("alfa"),
		give: []*Tree{NewFile("bravo"), NewDir("charlie")},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie", IsDir: true},
		}},
	}, {
		name: "has children",
		tree: NewDir("alfa").AddFiles("bravo"),
		give: []*Tree{NewFile("charlie"), NewDir("delta")},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie"},
			{Name: "delta", IsDir: true},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.Add(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAddDir(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Tree
		give      string
		wantTree  *Tree
		wantChild *Tree
	}{{
		name: "empty",
		tree: NewDir("alfa"),
		give: "bravo",
		wantTree: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo", IsDir: true},
		}},
		wantChild: &Tree{Name: "bravo", IsDir: true},
	}, {
		name: "has children",
		tree: NewDir("alfa").AddFiles("bravo"),
		give: "charlie",
		wantTree: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie", IsDir: true},
		}},
		wantChild: &Tree{Name: "charlie", IsDir: true},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChild := tt.tree.AddDir(tt.give)
			assert.DeepEqual(t, tt.tree, tt.wantTree, cmpOptions)
			assert.DeepEqual(t, gotChild, tt.wantChild, cmpOptions)
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
		tree: NewDir("alfa"),
		give: []string{"bravo", "charlie"},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo", IsDir: true},
			{Name: "charlie", IsDir: true},
		}},
	}, {
		name: "has children",
		tree: NewDir("alfa").AddFiles("bravo"),
		give: []string{"charlie", "delta"},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie", IsDir: true},
			{Name: "delta", IsDir: true},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.AddDirs(tt.give...)
			assert.DeepEqual(t, got, tt.want, cmpOptions)
		})
	}
}

func TestTreeAddFile(t *testing.T) {
	tests := []struct {
		name      string
		tree      *Tree
		give      string
		wantTree  *Tree
		wantChild *Tree
	}{{
		name: "empty",
		tree: NewDir("alfa"),
		give: "bravo",
		wantTree: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
		}},
		wantChild: &Tree{Name: "bravo"},
	}, {
		name: "has children",
		tree: NewDir("alfa").AddFiles("bravo"),
		give: "charlie",
		wantTree: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie"},
		}},
		wantChild: &Tree{Name: "charlie"},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChild := tt.tree.AddFile(tt.give)
			assert.DeepEqual(t, tt.tree, tt.wantTree, cmpOptions)
			assert.DeepEqual(t, gotChild, tt.wantChild, cmpOptions)
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
		tree: NewDir("alfa"),
		give: []string{"bravo", "charlie"},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie"},
		}},
	}, {
		name: "has children",
		tree: NewDir("alfa").AddFiles("bravo"),
		give: []string{"charlie", "delta"},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo"},
			{Name: "charlie"},
			{Name: "delta"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tree.AddFiles(tt.give...)
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
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo", IsDir: true, Children: []*Tree{
				{Name: "foxtrot", IsDir: true, Children: []*Tree{
					{Name: "india.txt"},
				}},
				{Name: "golf.txt"},
				{Name: "hotel", IsDir: true},
			}},
			{Name: "charlie.txt"},
			{Name: "delta", IsDir: true},
			{Name: "echo.txt"},
			{Name: "kilo", IsDir: true, Children: []*Tree{
				{Name: "juliet.txt"},
			}},
		}},
	}, {
		name: "directories before files",
		give: []SortOption{WithDirsFirst(true)},
		want: &Tree{Name: "alfa", IsDir: true, Children: []*Tree{
			{Name: "bravo", IsDir: true, Children: []*Tree{
				{Name: "foxtrot", IsDir: true, Children: []*Tree{
					{Name: "india.txt"},
				}},
				{Name: "hotel", IsDir: true},
				{Name: "golf.txt"},
			}},
			{Name: "delta", IsDir: true},
			{Name: "kilo", IsDir: true, Children: []*Tree{
				{Name: "juliet.txt"},
			}},
			{Name: "charlie.txt"},
			{Name: "echo.txt"},
		}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewDir("alfa").Add(
				NewFile("charlie.txt"),
				NewDir("bravo").Add(
					NewFile("golf.txt"),
					NewDir("foxtrot").AddFiles("india.txt"),
					NewDir("hotel"),
				),
				NewDir("kilo").AddFiles("juliet.txt"),
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
		tree: NewDir("alfa"),
		want: `alfa`,
	}, {
		name: "nested",
		tree: NewDir("alfa").Add(
			NewFile("bravo.txt"),
			NewDir("charlie").Add(
				NewFile("delta.txt"),
				NewDir("echo").AddFiles("foxtrot.txt", "golf.txt"),
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
		tree: NewDir("alfa").Add(
			NewDir("bravo").AddFiles("charlie.txt"),
			NewFile("delta.txt"),
		),
		want: `alfa
├── bravo
│   └── charlie.txt
└── delta.txt`,
	}, {
		name: "multiline names",
		tree: NewDir("alfa\n[dir]\n[3 MB]").AddFiles(
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
