package word

import (
	"bytes"
	"os/exec"
)

// function to get the user selected word from an active windows using xclip tool
func GetSelectedWord() (string, int) {
	cmd := exec.Command("xclip", "-o")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err.Error(), 1
	}
	res := out.String()
	return res, 0
}

func SendNotification(word string) {
	cmd := exec.Command("notify-send", word)
	cmd.Run()
}
