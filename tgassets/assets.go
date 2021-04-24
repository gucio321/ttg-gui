package tgassets

import (
	_ "embed"
)

var (
	//go:embed images/x.png
	ImageX []byte
	//go:embed images/o.png
	ImageO []byte
)
