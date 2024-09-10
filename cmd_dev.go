package yaakcli

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"os"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Build plugin bundle continuously when the filesystem changes",
	Long:  "Monitor the filesystem and build the plugin bundle when something changes. Useful for plugin development.",
	Run: func(cmd *cobra.Command, args []string) {
		if !fileExists("./package.json") {
			ExitError("./package.json does not exist. Ensure that you are in a plugin directory?")
		}

		srcPath := "./src/index.ts"
		ctx, errors := api.Context(ESLintBuildOptions([]string{srcPath}))
		if errors != nil {
			println("Failed to create esbuild context")
			os.Exit(1)
		}

		fmt.Printf("Watching %s...\n", srcPath)

		err := ctx.Watch(api.WatchOptions{})
		CheckError(err)

		// Returning from main() exits immediately in Go.
		// Block forever so that we keep watching and don't exit.
		<-make(chan struct{})
	},
}
