package asciitree

import "testing"

var result string

func BenchmarkNodeAdd(b *testing.B) {
	trees := []*Node{
		NewFile("bravo.txt"),
		NewFile("charlie.txt"),
		NewFile("delta.txt"),
		NewFile("echo.txt"),
		NewFile("foxtrot.txt"),
		NewFile("golf.txt"),
		NewFile("hotel.txt"),
		NewFile("india.txt"),
		NewFile("juliet.txt"),
		NewFile("kilo.txt"),
	}
	for i := 0; i < b.N; i++ {
		tree := NewDir("alfa")
		tree.Add(trees...)
	}
}

func BenchmarkAddDirs(b *testing.B) {
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
		tree := NewDir("alfa")
		tree.AddDirs(names...)
	}
}

func BenchmarkAddFiles(b *testing.B) {
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
		tree := NewDir("alfa")
		tree.AddFiles(names...)
	}
}

func BenchmarkNodeString(b *testing.B) {
	tree := NewDir("alfa").Add(
		NewFile("bravo.txt"),
		NewDir("charlie").Add(
			NewFile("delta.txt"),
			NewDir("echo").AddFiles("foxtrot.txt"),
			NewDir("golf").Add(
				NewDir("hotel\n[dir]\n[3 MB]").AddFiles(
					"india.txt\n[file]\n[1 MB]",
					"juliet.txt\n[file]\n[2 MB]",
				),
			),
		),
		NewDir("kilo").AddFiles("lima.txt"),
		NewDir("mike"),
		NewFile("november.txt"),
	)
	for i := 0; i < b.N; i++ {
		result = tree.String()
	}
}
