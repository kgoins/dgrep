package cmd

import (
	"fmt"
	"log"

	"github.com/kgoins/dgrep/dgrep"
	"github.com/kgoins/dgrep/internal"
	"github.com/spf13/cobra"
)

var ipsCmd = &cobra.Command{
	Use:   "ips filename",
	Short: "Extract IP addresses from input",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		isBin, _ := cmd.Flags().GetBool("binary")

		inStrs, err := internal.ReadStrings(filename, isBin)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		resultAddrs := dgrep.ExtractIPs(inStrs)

		publicOnly, _ := cmd.Flags().GetBool("public")
		if publicOnly {
			resultAddrs = dgrep.GetPublicIPs(resultAddrs)
		}

		for _, result := range resultAddrs {
			fmt.Println(result.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(ipsCmd)

	ipsCmd.Flags().BoolP(
		"public",
		"p",
		false,
		"Filter out all RFC 1918 IP addresses and return only public IPs",
	)
}
