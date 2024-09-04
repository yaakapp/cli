package yaakcli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "plaak",
	Short: "Develop plugins for Yaak",
	Long:  `Generate, build, and debug plugins for Yaak, the most intuitive desktop API client`,
	Run: func(cmd *cobra.Command, args []string) {
		println("Hello from Plaak")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(generateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
