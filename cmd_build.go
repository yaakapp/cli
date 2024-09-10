package yaakcli

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build entrypoint",
	Short: "Transpile code into a runnable plugin bundle",
	Run: func(cmd *cobra.Command, args []string) {
		if !fileExists("./package.json") {
			ExitError("./package.json does not exist. Ensure that you are in a plugin directory?")
		}

		srcPath := "./src/index.ts"
		fmt.Printf("Building %s...\n", srcPath)

		result := api.Build(ESLintBuildOptions([]string{srcPath}))
		for _, o := range result.OutputFiles {
			fmt.Printf("Compiled to: %s\n", o.Path)
		}
	},
}
