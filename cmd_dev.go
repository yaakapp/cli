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
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, errors := api.Context(ESLintBuildOptions(args))
		if errors != nil {
			println("Failed to create esbuild context")
			os.Exit(1)
		}

		err := ctx.Watch(api.WatchOptions{})
		CheckError(err)

		fmt.Printf("watching %s...\n", args[0])

		// Returning from main() exits immediately in Go.
		// Block forever so that we keep watching and don't exit.
		<-make(chan struct{})
	},
}
