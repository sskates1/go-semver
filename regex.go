package main

import (
	"regexp"
)

func getRegexes() map[string]*regexp.Regexp {
	regValues := make(map[string]string)

	// The following Regular Expressions can be used for tokenizing,
	// validating, and parsing SemVer version strings.

	// ## Numeric Identifier
	// A single `0`, or a non-zero digit followed by zero or more digits.
	// regValues["NUMERICIDENTIFIER"] = "0|[1-9]\\d*"
	// regValues["NUMERICIDENTIFIERLOOSE"] = "[0-9]+"
	regValues["NUMERICIDENTIFIER"] = "[0-9]+"

	// ## Non-numeric Identifier
	// Zero or more digits, followed by a letter or hyphen, and then zero or
	// more letters, digits, or hyphens.
	regValues["NONNUMERICIDENTIFIER"] = "\\d*[a-zA-Z-][a-zA-Z0-9-]*"

	// ## Main Version
	// Three dot-separated numeric identifiers.
	regValues["MAINVERSION"] = "(" + regValues["NUMERICIDENTIFIER"] + ")\\." +
		"(" + regValues["NUMERICIDENTIFIER"] + ")\\." +
		"(" + regValues["NUMERICIDENTIFIER"] + ")"

	// Star ranges basically just allow anything at all.
	regValues["STAR"] = "(<|>)?=?\\s*\\*"

	regexes := make(map[string]*regexp.Regexp)
	var err error

	for key, value := range regValues {
		regexes[key], err = regexp.Compile(value)
		panicIfError(err, key)
	}

	return regexes
}
