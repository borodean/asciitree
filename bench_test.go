package asciitree

import "testing"

var result string

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
