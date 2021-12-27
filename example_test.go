package asciitree_test

import (
	"fmt"

	"github.com/borodean/asciitree"
)

func ExampleNewDir() {
	dir := asciitree.NewDir("albums")
	fmt.Printf("%#v", dir)
	// Output:
	// &asciitree.Node{Name:"albums", IsDir:true, Children:[]*asciitree.Node(nil)}
}

func ExampleNewFile() {
	file := asciitree.NewFile("ONUKA.jpg")
	fmt.Printf("%#v", file)
	// Output:
	// &asciitree.Node{Name:"ONUKA.jpg", IsDir:false, Children:[]*asciitree.Node(nil)}
}

func ExampleNode_Add() {
	tree := asciitree.NewDir("albums")
	tree.Add(asciitree.NewDir("VIDLIK"), asciitree.NewFile("ONUKA.jpg"), asciitree.NewFile(".DS_Store"))
	fmt.Println(tree)
	// Output:
	// albums
	// ├── VIDLIK
	// ├── ONUKA.jpg
	// └── .DS_Store
}

func ExampleNode_AddDir() {
	tree := asciitree.NewDir("albums")
	dir := tree.AddDir("VIDLIK")
	fmt.Println(tree)
	fmt.Println("--")
	fmt.Println(dir)
	// Output:
	// albums
	// └── VIDLIK
	// --
	// VIDLIK
}

func ExampleNode_AddDirs() {
	tree := asciitree.NewDir("albums")
	tree.AddDirs("VIDLIK", "KOLIR")
	fmt.Println(tree)
	// Output:
	// albums
	// ├── VIDLIK
	// └── KOLIR
}

func ExampleNode_AddFile() {
	tree := asciitree.NewDir("albums")
	file := tree.AddDir("ONUKA.jpg")
	fmt.Println(tree)
	fmt.Println("--")
	fmt.Println(file)
	// Output:
	// albums
	// └── ONUKA.jpg
	// --
	// ONUKA.jpg
}

func ExampleNode_AddFiles() {
	tree := asciitree.NewDir("albums")
	tree.AddFiles(".DS_Store", "ONUKA.jpg")
	fmt.Println(tree)
	// Output:
	// albums
	// ├── .DS_Store
	// └── ONUKA.jpg
}

func ExampleNode_Sort() {
	tree := asciitree.NewDir("albums").Add(
		asciitree.NewFile(".DS_Store"),
		asciitree.NewFile("ONUKA.jpg"),
		asciitree.NewDir("VIDLIK").AddFiles(
			"Svitanok.mp3",
			"Vidlik.mp3",
			"Other (Intro).mp3",
			"Other.mp3",
			"19 86.mp3",
		),
		asciitree.NewDir("KOLIR").AddFiles(
			"VSTUP.mp3",
			"CEAHC.mp3",
			"ZENIT.mp3",
			"UYAVY (feat. DakhaBrakha).mp3",
			"TY.mp3",
			"GUMA.mp3",
			"XASHI.mp3",
			"NA SAMOTI.mp3",
			"SON.mp3",
			"23: 42.mp3",
		),
	)
	fmt.Println(tree.Sort())
	// Output:
	// albums
	// ├── .DS_Store
	// ├── KOLIR
	// │   ├── 23: 42.mp3
	// │   ├── CEAHC.mp3
	// │   ├── GUMA.mp3
	// │   ├── NA SAMOTI.mp3
	// │   ├── SON.mp3
	// │   ├── TY.mp3
	// │   ├── UYAVY (feat. DakhaBrakha).mp3
	// │   ├── VSTUP.mp3
	// │   ├── XASHI.mp3
	// │   └── ZENIT.mp3
	// ├── ONUKA.jpg
	// └── VIDLIK
	//     ├── 19 86.mp3
	//     ├── Other (Intro).mp3
	//     ├── Other.mp3
	//     ├── Svitanok.mp3
	//     └── Vidlik.mp3
}

// To place directories before files, enable the WithDirsFirst option.
func ExampleNode_Sort_withDirsFirst() {
	tree := asciitree.NewDir("albums").Add(
		asciitree.NewFile(".DS_Store"),
		asciitree.NewFile("ONUKA.jpg"),
		asciitree.NewDir("VIDLIK").AddFiles(
			"Svitanok.mp3",
			"Vidlik.mp3",
			"Other (Intro).mp3",
			"Other.mp3",
			"19 86.mp3",
		),
		asciitree.NewDir("KOLIR").AddFiles(
			"VSTUP.mp3",
			"CEAHC.mp3",
			"ZENIT.mp3",
			"UYAVY (feat. DakhaBrakha).mp3",
			"TY.mp3",
			"GUMA.mp3",
			"XASHI.mp3",
			"NA SAMOTI.mp3",
			"SON.mp3",
			"23: 42.mp3",
		),
	)
	fmt.Println(tree.Sort(asciitree.WithDirsFirst(true)))
	// Output:
	// albums
	// ├── KOLIR
	// │   ├── 23: 42.mp3
	// │   ├── CEAHC.mp3
	// │   ├── GUMA.mp3
	// │   ├── NA SAMOTI.mp3
	// │   ├── SON.mp3
	// │   ├── TY.mp3
	// │   ├── UYAVY (feat. DakhaBrakha).mp3
	// │   ├── VSTUP.mp3
	// │   ├── XASHI.mp3
	// │   └── ZENIT.mp3
	// ├── VIDLIK
	// │   ├── 19 86.mp3
	// │   ├── Other (Intro).mp3
	// │   ├── Other.mp3
	// │   ├── Svitanok.mp3
	// │   └── Vidlik.mp3
	// ├── .DS_Store
	// └── ONUKA.jpg
}

func ExampleNode_String() {
	tree := asciitree.NewDir("albums").Add(
		asciitree.NewFile(".DS_Store"),
		asciitree.NewFile("ONUKA.jpg"),
		asciitree.NewDir("VIDLIK").AddFiles(
			"Svitanok.mp3",
			"Vidlik.mp3",
			"Other (Intro).mp3",
			"Other.mp3",
			"19 86.mp3",
		),
		asciitree.NewDir("KOLIR").AddFiles(
			"VSTUP.mp3",
			"CEAHC.mp3",
			"ZENIT.mp3",
			"UYAVY (feat. DakhaBrakha).mp3",
			"TY.mp3",
			"GUMA.mp3",
			"XASHI.mp3",
			"NA SAMOTI.mp3",
			"SON.mp3",
			"23: 42.mp3",
		),
	)
	// asciitree.Node implements the fmt.Stringer interface; hence the String call
	// can be omitted here to yield the same result:
	// fmt.Println(tree)
	fmt.Println(tree.String())
	// Output:
	// albums
	// ├── .DS_Store
	// ├── ONUKA.jpg
	// ├── VIDLIK
	// │   ├── Svitanok.mp3
	// │   ├── Vidlik.mp3
	// │   ├── Other (Intro).mp3
	// │   ├── Other.mp3
	// │   └── 19 86.mp3
	// └── KOLIR
	//     ├── VSTUP.mp3
	//     ├── CEAHC.mp3
	//     ├── ZENIT.mp3
	//     ├── UYAVY (feat. DakhaBrakha).mp3
	//     ├── TY.mp3
	//     ├── GUMA.mp3
	//     ├── XASHI.mp3
	//     ├── NA SAMOTI.mp3
	//     ├── SON.mp3
	//     └── 23: 42.mp3
}
