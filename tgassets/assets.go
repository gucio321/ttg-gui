package tgassets

import (
	_ "embed" // enables go:embed directive
)

// there are variables describing game images
// nolint:gochecknoglobals // go:embed - requires global variable
var (
	//go:embed images/x.png
	ImageX []byte
	//go:embed images/o.png
	ImageO []byte
)
