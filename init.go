package main

import (
	"os"
	"os/exec"
)

func main() {
	command := exec.Command("redis-server")
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		panic(err)
	}
}
