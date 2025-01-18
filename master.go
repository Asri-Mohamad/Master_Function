package Master_Function

import (
	"os"
	"os/exec"

	"github.com/fatih/color"
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
func SetColor(name string) func(...interface{}) string {
	var c color.Attribute
	switch name {
	case "blue", "Blue":
		c = color.FgBlue
	case "read", "Read":
		c = color.FgRed
	case "green", "Green":
		c = color.FgGreen
	case "yellow", "Yellow":
		c = color.FgYellow
	case "pinc", "Pinc":
		c = color.FgHiMagenta

	default:
		c = color.FgWhite
	}

	return color.New(c).SprintFunc()

}
