// +build prod

package versiontags

import (
	_ "embed"
)

//go:embed version.txt
var version string