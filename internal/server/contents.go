/*
Copyright © 2023 Ajith

*/
package server

type Property struct {
	Name		string
	Path		string
	Size		string
}

type Content struct {
	Path			string
	Files			[]Property
	Directories		[]Property
}

func NewProperty() *Property {
	return &Property{
		Name: 	"",
		Path:	"",
		Size:	"",
	}
}

func (p *Property) SetName(name string) {
	p.Name = name
}

func (p *Property) SetPath(path string) {
	p.Path = path
}

func (p *Property) SetSize(size string) {
	p.Size = size
}

func NewContent() *Content {
	return &Content{
		Path:			"",
		Files:			[]Property{},
		Directories:	[]Property{},
	}
}

func (c *Content) SetPath(path string) {
	c.Path = path
}

func (c *Content) addFile(file *Property) {
	c.Files = append(c.Files, *file)
}

func (c *Content) addDirectory(directory *Property) {
	c.Directories = append(c.Directories, *directory)
}
