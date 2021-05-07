package version

import (
	_ "embed"
	"fmt"
)

//go:embed version.txt
var version string

//go:embed version2.txt
var version2 string

func main() {
	fmt.Println(version)
	fmt.Println(version2)
}