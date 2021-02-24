package internal

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func extractStrings(filename string) ([]string, error) {
	stringsCmd := exec.Command("strings", filename)

	err := stringsCmd.Run()
	if err != nil {
		return nil, err
	}

	rawOutput, err := stringsCmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(rawOutput), "\n"), nil
}

// ReadStrings will run the strings cmd to extract all printable strings from a binary or text file.
func ReadStrings(filename string, isBin bool) ([]string, error) {
	var results []string
	var err error

	filename = strings.TrimSpace(filename)
	filename, err = homedir.Expand(filename)
	if err != nil {
		return nil, err
	}

	fullPath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	if isBin {
		results, err = extractStrings(fullPath)
		if err != nil {
			return nil, err
		}

		return results, nil
	}

	fileBytes, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return []string{string(fileBytes)}, nil
}
