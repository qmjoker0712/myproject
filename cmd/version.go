package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"myproject/version"
)

var (
	// VersionFlag version Flag
	VersionFlag *bool
)

// VersionCmd version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the app version",
	Long: `usage example:
	api(.exe) version
	print the app version`,
	Run: func(cmd *cobra.Command, args []string) {
		if *VersionFlag {
			v := version.Get()
			marshalled, err := json.MarshalIndent(&v, "", "  ")
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}

			fmt.Println(string(marshalled))
		}

	},
}
