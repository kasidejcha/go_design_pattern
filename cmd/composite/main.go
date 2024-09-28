package main

import "fmt"

// Component interface
type Component interface {
	Display(indent string)
}

// Leaf: File
type File struct {
	name string
}

func (f *File) Display(indent string) {
	fmt.Println(indent + f.name)
}

// Composite: Directory
type Directory struct {
	name       string
	components []Component
}

func (d *Directory) Add(component Component) {
	d.components = append(d.components, component)
}

func (d *Directory) Display(indent string) {
	fmt.Println(indent + d.name)
	for _, component := range d.components {
		component.Display(indent + "  ")
	}
}

func main() {
	// Create files
	file1 := &File{name: "file1.txt"}
	file2 := &File{name: "file2.txt"}
	file3 := &File{name: "file3.txt"}

	// Create directories
	rootDir := &Directory{name: "root"}
	subDir1 := &Directory{name: "subdir1"}
	subDir2 := &Directory{name: "subdir2"}

	// Build the tree structure
	rootDir.Add(file1)
	rootDir.Add(subDir1)

	subDir1.Add(file2)
	subDir1.Add(subDir2)

	subDir2.Add(file3)

	// Display the structure
	startingIndent := ""
	rootDir.Display(startingIndent)
}
