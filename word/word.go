package word

import (
	"bytes"
	"os/exec"
)

func GetSelectedWord() string {
	cmd := exec.Command("xclip", "-o")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		result := "failed"
		return result
	}
	res := out.String()
	return res
}
