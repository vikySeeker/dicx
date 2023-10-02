package word

import (
	"bytes"
	"os/exec"
	"strings"
)

// function to get the user selected word from an active windows using xclip tool
func GetSelectedWord() (string, error) {
	cmd := exec.Command("xclip", "-o")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "Error occured", err
	}
	res := out.String()
	res = strings.TrimSpace(res)
	res = strings.Split(res, " ")[0]
	res = strings.ReplaceAll(res, ".", "")
	return res, nil
}
