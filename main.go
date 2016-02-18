package main

import (
	"fmt"
)

func main() {
	inClValues := inSemVer()

	semver, err := parse(inClValues.firstVersion)
	panicIfError(err, "parse failed")

	fmt.Println(*semver.Version())
}
