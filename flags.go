package main

import "flag"

type inCLValues struct {
	firstVersion string
}

func inSemVer() *inCLValues {
	out := new(inCLValues)
	firstString := flag.String("s", "9.9.9", "The semver string you want to parse")

	flag.Parse()
	out.firstVersion = *firstString

	return out
}
