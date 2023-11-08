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

func NewContent() *Content {
	return &Content{
		Path:			"",
		Files:			[]Property{},
		Directories:	[]Property{},
	}
}
