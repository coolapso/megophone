package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "DEV"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print megophone version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Megophone %s\n", Version)
	},
}
