[![Go Reference][go-reference-badge]][go-reference]
[![Codacy Quality][codacy-quality-badge]][codacy-dashboard]
[![Codecov Coverage][codecov-coverage-badge]][codecov-coverage]

# ![asciitree][asciitree-logo]

Human-friendly Go module that builds and prints directory trees using ASCII art.

## Installation

```bash
go get github.com/borodean/asciitree
```

## Usage

```go
// Build a directory tree representation:
tree := asciitree.NewDir("albums").Add(
  asciitree.NewFile("ONUKA.jpg"),
  asciitree.NewDir("VIDLIK").AddFiles(
    "Svitanok.mp3",
    "Vidlik.mp3",
  ),
  asciitree.NewDir("KOLIR").AddFiles(
    "CEAHC.mp3",
    "ZENIT.mp3",
    "UYAVY (feat. DakhaBrakha).mp3",
    "XASHI.mp3",
  ),
)

// Sort the tree's descendants alphanumerically while placing directories
// before files:
tree.Sort(asciitree.WithDirsFirst(true))

// Print an ASCII art representation of the directory tree:
fmt.Println(tree)
```

## License

MIT.

[go-reference-badge]: https://pkg.go.dev/badge/github.com/borodean/asciitree.svg
[go-reference]: https://pkg.go.dev/github.com/borodean/asciitree
[codacy-quality-badge]: https://app.codacy.com/project/badge/Grade/c5ef187cb0fa41f4ad4fa4f635cc8cd6
[codacy-dashboard]: https://www.codacy.com/gh/borodean/asciitree/dashboard
[codecov-coverage-badge]: https://codecov.io/gh/borodean/asciitree/branch/main/graph/badge.svg?token=UNHXM8TZJ4
[codecov-coverage]: https://codecov.io/gh/borodean/asciitree
[asciitree-logo]: ./logo.svg
