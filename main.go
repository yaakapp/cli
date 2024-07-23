package main

import (
	"flag"
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"os"
)

func main() {
	flag.Parse()

	options := api.BuildOptions{
		EntryPoints: flag.Args(),
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
		fmt.Fprintf(os.Stderr, "Compiled to: %s\n", o.Path)
	}
}
