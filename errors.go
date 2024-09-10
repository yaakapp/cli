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
	ExitError(err.Error())
}

func ExitError(msg string) {
	pterm.Println(pterm.Red(fmt.Sprintf("Error: %s", msg)))
	os.Exit(1)
}
