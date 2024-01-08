package main_test

import (
	"os/exec"
	"slices"
	"testing"
)

func executeReq(t *testing.T, comm string, param []string) string {
	cmd := exec.Command(comm, param...)

	res, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	return string(res)
}

func TestSort(t *testing.T) {
	filepath := "sample"
	parameters := [][]string{
		{"-k", "4", "-n", filepath},
		{"-k", "7", "-r", "-u", filepath},
		{"-k", "2", filepath},
	}

	for _, param := range parameters {
		sysOut := executeReq(t, "sort", param)
		programOut := executeReq(t, "go", slices.Insert(param, 0, "run", "task.go"))
		if sysOut != programOut {
			t.Errorf(
				"answer wrong, param: %s\nsystem result:\n------\n%s------\nprogram result:\n------\n%s\n",
				param,
				sysOut,
				programOut,
			)
		}
	}
}
