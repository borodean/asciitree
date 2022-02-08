package asciitree_test

import (
	"testing"

	"github.com/borodean/asciitree"
)

var result string

func BenchmarkNode_Add(b *testing.B) {
	trees := []*asciitree.Node{
		asciitree.NewFile("bravo.txt"),
		asciitree.NewFile("charlie.txt"),
		asciitree.NewFile("delta.txt"),
		asciitree.NewFile("echo.txt"),
		asciitree.NewFile("foxtrot.txt"),
		asciitree.NewFile("golf.txt"),
		asciitree.NewFile("hotel.txt"),
		asciitree.NewFile("india.txt"),
		asciitree.NewFile("juliet.txt"),
		asciitree.NewFile("kilo.txt"),
	}

	for i := 0; i < b.N; i++ {
		tree := asciitree.NewDir("alfa")
		tree.Add(trees...)
	}
}

func BenchmarkNode_AddDirs(b *testing.B) {
	names := []string{
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
		"india",
		"juliet",
		"kilo",
	}

	for i := 0; i < b.N; i++ {
		tree := asciitree.NewDir("alfa")
		tree.AddDirs(names...)
	}
}

func BenchmarkNode_AddFiles(b *testing.B) {
	names := []string{
		"bravo.txt",
		"charlie.txt",
		"delta.txt",
		"echo.txt",
		"foxtrot.txt",
		"golf.txt",
		"hotel.txt",
		"india.txt",
		"juliet.txt",
		"kilo.txt",
	}

	for i := 0; i < b.N; i++ {
		tree := asciitree.NewDir("alfa")
		tree.AddFiles(names...)
	}
}

func BenchmarkNode_Sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := asciitree.NewDir("juliet").Add(
			asciitree.NewFile("golf.txt"),
			asciitree.NewDir("foxtrot").Add(
				asciitree.NewFile("charlie.txt"),
				asciitree.NewDir("mike").AddFiles("lima.txt"),
				asciitree.NewDir("india").Add(
					asciitree.NewDir("delta\n[dir]\n[3 MB]").AddFiles(
						"november.txt\n[file]\n[1 MB]",
						"echo.txt\n[file]\n[2 MB]",
					),
				),
			),
			asciitree.NewDir("bravo").AddFiles("kilo.txt"),
			asciitree.NewDir("alfa"),
			asciitree.NewFile("hotel.txt"),
		)
		tree.Sort()
	}
}

func BenchmarkNode_String(b *testing.B) {
	tree := asciitree.NewDir("alfa").Add(
		asciitree.NewFile("bravo.txt"),
		asciitree.NewDir("charlie").Add(
			asciitree.NewFile("delta.txt"),
			asciitree.NewDir("echo").AddFiles("foxtrot.txt"),
			asciitree.NewDir("golf").Add(
				asciitree.NewDir("hotel\n[dir]\n[3 MB]").AddFiles(
					"india.txt\n[file]\n[1 MB]",
					"juliet.txt\n[file]\n[2 MB]",
				),
			),
		),
		asciitree.NewDir("kilo").AddFiles("lima.txt"),
		asciitree.NewDir("mike"),
		asciitree.NewFile("november.txt"),
	)
	for i := 0; i < b.N; i++ {
		result = tree.String()
	}
}
