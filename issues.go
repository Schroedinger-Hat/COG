package main

import (
	"fmt"
	"math/rand"
)

var issues = []string{
	"vegeutils",
	"libgardening",
	"currykit",
	"spicerack",
	"fullenglish",
	"eggy",
	"bad-kitty",
	"chai",
	"hojicha",
	"libtacos",
	"babys-monads",
	"libpurring",
	"currywurst-devel",
	"xmodmeow",
	"licorice-utils",
	"cashew-apple",
	"rock-lobster",
	"standmixer",
	"coffee-CUPS",
	"libesszet",
	"zeichenorientierte-benutzerschnittstellen",
	"schnurrkit",
	"old-socks-devel",
	"jalapeño",
	"molasses-utils",
	"xkohlrabi",
	"party-gherkin",
	"snow-peas",
	"libyuzu",
}

func GetIssue() []string {
	workingIssues := issues
	copy(workingIssues, issues)

	rand.Shuffle(len(workingIssues), func(i, j int) {
		workingIssues[i], workingIssues[j] = workingIssues[j], workingIssues[i]
	})

	for k := range workingIssues {
		workingIssues[k] += fmt.Sprintf("-%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10)) //nolint:gosec
	}
	return workingIssues
}
