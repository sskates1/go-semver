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

	fmt.Println(version)
	semver, err := newSemVer(version)
	if err != nil {
		return nil, err
	}
	return semver, nil
}

func newSemVer(version string) (*SemVer, error) {
	semver := new(SemVer)
	semver.raw = strings.TrimSpace(version)

	regexs := getRegexes()
	var regex *regexp.Regexp
	regex = regexs["MAINVERSION"]

	matched := regex.FindAllString(semver.raw, -1)
	errors := make([]error, 3)
	semver.major, errors[0] = strconv.Atoi(matched[0])
	semver.minor, errors[1] = strconv.Atoi(matched[1])
	semver.patch, errors[2] = strconv.Atoi(matched[2])
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
