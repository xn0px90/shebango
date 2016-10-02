package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	Run("echo hello bash !")
}

func Run(command string) {
	fmt.Printf("run: %s\n", command)
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		panic(err)
	}
}
