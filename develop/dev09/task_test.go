package main_test

import (
	"os/exec"
	"testing"
)

func TestWget(t *testing.T) {
	url := "https://ru.wikipedia.org/wiki/Go"
	cmd := exec.Command("go", "run", "task.go", url)
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
}
