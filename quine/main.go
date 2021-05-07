package main

import (
	_ "embed"
	"fmt"
)

//go:embed version/version.txt
var src string

func main() {
	fmt.Print(src)
}