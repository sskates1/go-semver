package main

import "flag"

// InCLValues string from args
type InCLValues struct {
	firstVersion string
}

func inSemVer() *InCLValues {
	out := new(InCLValues)
	firstString := flag.String("s", "9.9.9", "The semver string you want to parse")

	flag.Parse()
	out.firstVersion = *firstString

	return out
}
