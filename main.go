package main

import (
	"flag"
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"os"
)

func main() {
	flag.Parse()

	result := api.Build(api.BuildOptions{
		EntryPoints: flag.Args(),
		Outfile:     "build/index.js",
		Bundle:      true, // Inline dependencies
		Write:       true, // Write to disk
		Format:      api.FormatCommonJS,
		LogLevel:    api.LogLevelInfo,
	})

	for _, o := range result.OutputFiles {
		fmt.Printf("Compiled to: %s\n", o.Path)
		fmt.Fprintf(os.Stderr, "Compiled to: %s\n", o.Path)
	}
}
