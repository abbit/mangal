package meta

import _ "embed"

//go:embed logo.txt
var Logo string

const (
	AppName = "mangal"
	Version = "5.4.5"
)
