package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BlockscapeLab/peerdx/addrbook"

	"github.com/spf13/cobra"
)

func createAddrbookCmd() *cobra.Command {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	dir := new(string)
	addrbookCmd := &cobra.Command{
		Use:   "addrbook",
		Short: "compare address books",
		Long:  `compares address books that are local on disk in provided directory`,

		Run: func(cmd *cobra.Command, args []string) {
			loadConfig()
			addrbook.LoadAndCompare(*dir)
		},
	}
	addrbookCmd.Flags().StringVarP(dir, "dir", "d", filepath.Join(workingDir, "addrBooks"),
		"the directory containing the address books.'")

	return addrbookCmd
}
