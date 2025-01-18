package Master_Function

import (
	"fmt"
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
func CharGetKey() byte {

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

// ------------------------ Set color ------------------------------
func ColorText(c string, text string) string {
	var colorCodeGet string
	switch c {
	case "Black", "black":
		colorCodeGet = "30"
	case "read", "Read":
		colorCodeGet = "31"
	case "green", "Green":
		colorCodeGet = "32"
	case "yellow", "Yellow":
		colorCodeGet = "33"
	case "blue", "Blue":
		colorCodeGet = "34"
	case "Magenta", "magenta":
		colorCodeGet = "35"
	case "cyan", "Cyan":
		colorCodeGet = "36"
	case "Reset", "reset":
		colorCodeGet = "0"

	default:
		colorCodeGet = "0"
	}

	return fmt.Sprintf("\033[%sm%s\033[0m", colorCodeGet, text)

}
