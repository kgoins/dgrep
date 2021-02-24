package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kgoins/dgrep/urlparser"
	hashset "github.com/kgoins/hashset/pkg"
	"github.com/spf13/cobra"
)

var etldCmd = &cobra.Command{
	Use:   "etld filename",
	Short: "Extract all effective TLDs from input domains",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		filename = strings.TrimSpace(filename)
		unique, _ := cmd.Flags().GetBool("unique")

		etlds, err := getETLDsFromFile(filename, unique)
		if err != nil {
			log.Fatalf(err.Error())
		}

		for _, etld := range etlds {
			fmt.Println(etld)
		}
	},
}

func getETLDsFromFile(filename string, unique bool) ([]string, error) {
	domainFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer domainFile.Close()

	results := make([]string, 0, 0)
	scanner := bufio.NewScanner(domainFile)

	for scanner.Scan() {
		lineText := scanner.Text()
		parsedTLD, err := urlparser.Parse(lineText)
		if err != nil {
			continue
		}

		etld := parsedTLD.Domain + "." + parsedTLD.TLD
		results = append(results, etld)
	}

	if unique {
		set := hashset.NewStrHashset(results...)
		results = set.Values()
	}

	return results, nil
}

func init() {
	rootCmd.AddCommand(etldCmd)

	etldCmd.Flags().BoolP(
		"unique",
		"u",
		false,
		"Return only unique values",
	)
}
