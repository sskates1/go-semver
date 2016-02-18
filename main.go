package main

import (
	"fmt"
)

func main() {
	_ = getRegexes()
	fmt.Println("regexes loaded properly")

	semver, err := parse("1.2.3")
	panicIfError(err, "parse failed")

	fmt.Println(semver.version())
}
