package cmd

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var prependCmd = &cobra.Command{
	Use:   "prepend domainFile wordlist",
	Short: "Prepend will prepend each entry in the wordlist to each domain in the input file",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domainsFilePath := args[0]
		wordlistPath := args[1]

		wordlist, err := os.Open(wordlistPath)
		if err != nil {
			log.Fatalf("Unable to open wordlist: %s", err.Error())
		}
		defer wordlist.Close()

		domainsFile, err := os.Open(domainsFilePath)
		if err != nil {
			log.Fatalf("Unable to open domains file: %s", err.Error())
		}
		defer domainsFile.Close()

		var writer io.WriteCloser
		outFile, _ := cmd.Flags().GetString("outfile")
		if outFile != "" {
			writer, err = os.OpenFile(
				outFile,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				0644,
			)
			if err != nil {
				log.Fatalf("Unable to open output file: %s", err.Error())
			}
		} else {
			writer = os.Stdout
		}
		defer writer.Close()

		err = prependWords(domainsFile, wordlist, writer)
		if err != nil {
			log.Fatalf("Error writing output: %s", err.Error())
		}
	},
}

func prependWords(domainSrc io.Reader, wordlistSrc io.ReadSeeker, output io.Writer) (err error) {
	domainsBuf := bufio.NewScanner(domainSrc)

	for domainsBuf.Scan() {
		err = prependToDomain(domainsBuf.Text(), wordlistSrc, output)
		if err != nil {
			return
		}

		_, err = wordlistSrc.Seek(0, io.SeekStart)
		if err != nil {
			return
		}
	}

	return nil
}

func prependToDomain(domain string, wordlistSrc io.Reader, output io.Writer) error {
	wordBuf := bufio.NewScanner(wordlistSrc)
	for wordBuf.Scan() {
		word := wordBuf.Text()
		if word == "" {
			continue
		}

		newDomain := word + "." + domain
		_, err := output.Write([]byte(newDomain + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(prependCmd)

	prependCmd.Flags().StringP(
		"outfile", "o", "", "output file",
	)
}
