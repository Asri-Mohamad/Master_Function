package Master_Function

import (
	"os"
	"os/exec"

	"golang.org/x/term"
)

// --------------- Clear Screen Function -------------------------
func Cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// -------------- Get a Charakter forom keybord -----------------
func CharakterheckKey() byte {

	// فعال کردن حالت خام (Raw Mode)
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {

		return 'e'
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState) // بازگرداندن تنظیمات اولیه

	// خواندن ورودی یک کاراکتر
	var key [1]byte
	_, err = os.Stdin.Read(key[:])
	if err != nil {

		return 'e'
	}
	return key[0]
}
