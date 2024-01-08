package main_test

import (
	"os/exec"
	"testing"
)

func TestTelnet(t *testing.T) {
	host := "google.com"
	port := "80"
	cmd := exec.Command("go", "run", "task.go", host, port)
	if err := cmd.Run(); err != nil {
		t.Error(err)
	}
}
