package yaakcli

import (
	"os"
	"path/filepath"
	"strings"
)

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

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
