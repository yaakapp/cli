package yaakcli

import (
	"github.com/spf13/cobra"
	"os"
)

func rootCmd(version string) *cobra.Command {
	var fVersion bool
	cmd := &cobra.Command{
		Use:   "yaakcli",
		Short: "Develop plugins for Yaak",
		Long:  `Generate, build, and debug plugins for Yaak, the most intuitive desktop API client`,
		Run: func(cmd *cobra.Command, args []string) {
			if fVersion {
				println(version)
				os.Exit(0)
			}

			checkErr(cmd.Help())
		},
	}
	cmd.AddCommand(buildCmd)
	cmd.AddCommand(generateCmd)

	cmd.Flags().BoolVar(&fVersion, "version", false, "Source directory to read from")

	return cmd
}

func Execute(version string) {
	checkErr(rootCmd(version).Execute())
}
