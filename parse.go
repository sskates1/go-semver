package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	//MAXLENGTH max size of semvar
	MAXLENGTH = 256
	//MaxUint64 used to calculate max int
	MaxUint64 = 1<<64 - 1
)

func parse(version string) (*SemVer, error) {
	if len(version) > MAXLENGTH {
		errorMessage := fmt.Sprintf("Version string longer than %d max length", MAXLENGTH)
		err := errors.New(errorMessage)
		return nil, err
	}
	if !testValid(version) {
		err := errors.New("Invalid semver format")
		return nil, err
	}

	semver, err := newSemVer(version)
	if err != nil {
		return nil, err
	}
	return semver, nil
}

func newSemVer(version string) (*SemVer, error) {
	semver := new(SemVer)
	semver.raw = strings.TrimSpace(version)
	semver.base = removeLeadingV(version)

	regexs := getRegexes()
	var regex *regexp.Regexp
	regex = regexs["MAINVERSION"]

	matched := regex.FindAllString(semver.base, -1)

	values := strings.Split(matched[0], ".")
	errors := make([]error, 3)
	semver.major, errors[0] = strconv.Atoi(values[0])
	semver.minor, errors[1] = strconv.Atoi(values[1])
	semver.patch, errors[2] = strconv.Atoi(values[2])
	for k, err := range errors {
		if err != nil {
			fmt.Println(k)
			panicIfError(err, "failed string to int conversion")
		}
	}

	return semver, nil
}

func testValid(version string) bool {

	return true
}

func removeLeadingV(version string) string {
	trimmedVersion := strings.TrimSpace(version)
	if strings.Index(trimmedVersion, "v") == 0 {
		trimmedVersion = strings.Replace(trimmedVersion, "v", "", 1)
	}
	return trimmedVersion
}
