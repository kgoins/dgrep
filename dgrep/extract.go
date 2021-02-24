package dgrep

import (
	"regexp"

	hashset "github.com/kgoins/hashset/pkg"
)

func extractWithRegex(inStrs []string, regex *regexp.Regexp) []string {
	results := hashset.NewStrHashset()

	for _, input := range inStrs {
		extracted := regex.FindAllString(input, -1)
		results.Add(extracted...)
	}

	return results.Values()
}
