package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the terminal screen across Windows, macOS, and Linux.
func ClearScreen() error {
	switch runtime.GOOS {
	case "windows":
		// Modern Windows Terminal and PowerShell support ANSI escape codes,
		// but executing "cls" via cmd ensures compatibility with legacy conhost.
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	default:
		// ANSI escape sequence:
		// \033[H  -> moves cursor to top-left (home)
		// \033[2J -> clears the entire visible screen
		// \033[3J -> clears the scrollback buffer (optional, remove if you want to keep history)
		_, err := fmt.Fprint(os.Stdout, "\033[H\033[2J\033[3J")
		return err
	}
}
