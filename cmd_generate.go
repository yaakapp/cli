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
	Use: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		defaultName := RandomName()
		defaultPath := "./" + defaultName

		_ = defaultPath

		dst, err := pterm.DefaultInteractiveTextInput.WithDefaultValue(defaultPath).Show()
		checkErr(err)

		if dirExists(dst) {
			returnError("")
		}

		pterm.Println("Generating plugin to:", pterm.Magenta(dst))

		// Create destination directory
		checkErr(os.MkdirAll(dst, 0755))

		// Copy static files
		copyFile("package.json", dst, defaultName)
		copyFile("tsconfig.json", dst, defaultName)
		copyFile("src/index.ts", dst, defaultName)

		c := exec.Command("npm", "install")
		c.Dir = dst
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		checkErr(c.Start())
		checkErr(c.Wait())
	},
}

func writeFile(path, contents string) {
	checkErr(os.MkdirAll(filepath.Dir(path), 0755))
	checkErr(os.WriteFile(path, []byte(contents), 0755))
}

func readFile(path string) string {
	pkgBytes, err := TemplateFS.ReadFile(path)
	checkErr(err)
	return string(pkgBytes)
}

func copyFile(relPath, dstDir, name string) {
	contents := readFile(filepath.Join("template", relPath))
	contents = strings.ReplaceAll(contents, "yaak-plugin-name", name)
	writeFile(filepath.Join(dstDir, relPath), contents)
}

func checkErr(err error) {
	if err == nil {
		return
	}

	pterm.Println(pterm.Red("Error: ", err.Error()))
	os.Exit(1)
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