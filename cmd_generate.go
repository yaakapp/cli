package yaakcli

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: `Generate a "Hello World" Yaak plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultName := RandomName()
		defaultPath := "./" + defaultName

		_ = defaultPath

		pluginDir, err := pterm.DefaultInteractiveTextInput.WithDefaultValue(defaultPath).Show()
		CheckError(err)

		if dirExists(pluginDir) {
			returnError("")
		}

		pterm.Println("Generating plugin to:", pterm.Magenta(pluginDir))

		// Create destination directory
		CheckError(os.MkdirAll(pluginDir, 0755))

		// Copy static files
		copyFile("package.json", pluginDir, defaultName)
		copyFile("tsconfig.json", pluginDir, defaultName)
		copyFile("src/index.ts", pluginDir, defaultName)
		copyFile("src/index.test.ts", pluginDir, defaultName)

		primary := pterm.NewStyle(pterm.FgLightWhite, pterm.BgMagenta, pterm.Bold)

		pterm.DefaultHeader.WithBackgroundStyle(primary).Println("Installing npm dependencies...")
		runCmd(pluginDir, "npm", "install")
		runCmd(pluginDir, "npm", "install", "@yaakapp/api")
		runCmd(pluginDir, "npm", "install", "-D", "@yaakapp/cli")
		runCmd(pluginDir, "npm", "run", "build")
	},
}

func runCmd(dir, cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	CheckError(c.Start())
	CheckError(c.Wait())
}

func writeFile(path, contents string) {
	CheckError(os.MkdirAll(filepath.Dir(path), 0755))
	CheckError(os.WriteFile(path, []byte(contents), 0755))
}

func readFile(path string) string {
	pkgBytes, err := TemplateFS.ReadFile(path)
	CheckError(err)
	return string(pkgBytes)
}

func copyFile(relPath, dstDir, name string) {
	contents := readFile(filepath.Join("template", relPath))
	contents = strings.ReplaceAll(contents, "yaak-plugin-name", name)
	writeFile(filepath.Join(dstDir, relPath), contents)
}

func dirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func returnError(msg string) {
	pterm.Println(pterm.Red(msg))
}
