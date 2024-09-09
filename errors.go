package yaakcli

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
)

func CheckError(err error) {
	if err == nil {
		return
	}

	pterm.Println(pterm.Red(fmt.Sprintf("Error: %s", err.Error())))
	os.Exit(1)
}
