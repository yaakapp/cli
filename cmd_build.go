package yaakcli

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Transpile code into a runnable plugin bundle",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		result := api.Build(ESLintBuildOptions(args))
		for _, o := range result.OutputFiles {
			fmt.Printf("Compiled to: %s\n", o.Path)
		}
	},
}
