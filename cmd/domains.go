package cmd

import (
	"fmt"
	"log"

	"github.com/kgoins/dgrep/dgrep"
	"github.com/kgoins/dgrep/internal"
	"github.com/spf13/cobra"
)

var domainsCmd = &cobra.Command{
	Use:   "domains filename",
	Short: "Extract domain strings from input",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		isBin, _ := cmd.Flags().GetBool("binary")

		inStrs, err := internal.ReadStrings(filename, isBin)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		results := dgrep.ExtractDomains(inStrs)

		for _, result := range results {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(domainsCmd)
}
