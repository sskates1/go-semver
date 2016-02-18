package main

import (
	"fmt"
)

//SemVer struct
type SemVer struct {
	raw        string
	base       string
	major      int
	minor      int
	patch      int
	preRelease string
	metadata   string
	loose      string
}

func (s SemVer) Version() *string {
	version := fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)
	return &version
}
