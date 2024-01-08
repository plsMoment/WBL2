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

func TestGrep(t *testing.T) {
	pattern := "needle"
	regexpPattern := "Nee.le"
	filepath := "sample"
	parameters := [][]string{
		{"-F", "-v", pattern, filepath},
		{"-F", "-i", "-c", pattern, filepath},
		{"-A", "2", "-B", "5", "-n", regexpPattern, filepath},
		{"-A", "2", "-i", "-n", regexpPattern, filepath},
	}

	for _, param := range parameters {
		sysOut := executeReq(t, "grep", param)
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
