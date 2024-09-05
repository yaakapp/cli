package main

import (
	"github.com/yaakapp/yaakcli"
)

var version = "dev"

func main() {
	yaakcli.Execute(version)
}
