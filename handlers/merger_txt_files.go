package handlers

import (
	"errors"
	"os/exec"
)

func mergeTxtFiles(args ...func() string) (string, error) {
	if len(args[0]())+len(args[1]()) < 6 {
		return "", errors.New("проверте пути к файлам")
	}
	cmd := exec.Command("python3", "scripts/merge.py", args[0](), args[1]())
	out, err := cmd.CombinedOutput()
	return string(out), err
}
