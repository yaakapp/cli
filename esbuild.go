package yaakcli

import (
	"github.com/evanw/esbuild/pkg/api"
)

func ESLintBuildOptions(entrypoints []string) api.BuildOptions {
	return api.BuildOptions{
		EntryPoints: entrypoints,
		Outfile:     "build/index.js",
		Platform:    api.PlatformNode,
		Bundle:      true, // Inline dependencies
		Write:       true, // Write to disk
		Format:      api.FormatCommonJS,
		LogLevel:    api.LogLevelInfo,
	}
}
