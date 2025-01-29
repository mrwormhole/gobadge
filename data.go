package main

// Replace with your data by using -ldflags like this:
//
// tinygo flash -target gobadge -ldflags="-X main.Name=@myontwitter -X main.Slogan1='Amazing human' -X main.Slogan2='also kind'"
//
// See Taskfile for more info.
var (
	Name, Slogan1, Slogan2 string
)

const (
	DefaultName    = "@TinyGolang"
	DefaultSlogan1 = "Go Compiler"
	DefaultSlogan2 = "Small Places"
)
