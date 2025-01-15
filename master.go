package master

import (
	"os"
	"os/exec"
)

func Cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
