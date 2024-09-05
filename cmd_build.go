package yaakcli

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/spf13/cobra"
	"os"
)

var buildCmd = &cobra.Command{
	Use:  "build",
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		options := api.BuildOptions{
			EntryPoints: []string{args[0]},
			Outfile:     "build/index.js",
			Platform:    api.PlatformNode,
			Bundle:      true, // Inline dependencies
			Write:       true, // Write to disk
			Format:      api.FormatCommonJS,
			LogLevel:    api.LogLevelInfo,
		}

		if os.Getenv("BUILD_PLATFORM") == "browser" {
			println("Set build platform to browser")
			options.Platform = api.PlatformBrowser
		} else {
			println("Set build platform to node")
		}

		result := api.Build(options)
		for _, o := range result.OutputFiles {
			fmt.Printf("Compiled to: %s\n", o.Path)
		}
	},
}
