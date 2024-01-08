package main_test

import (
	"os/exec"
	"slices"
	"testing"
)

func executeReq(t *testing.T, comm string, req []byte, param []string) string {
	cmd := exec.Command(comm, param...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Error(err)
	}

	go func() {
		stdin.Write(req)
		stdin.Close()
	}()

	res, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	return string(res)
}

func TestCut(t *testing.T) {
	requests := [][]byte{
		[]byte("a\tb\tcbvb\nabcfd\n1223112\t11\n"),
		[]byte("a:b:cbvb\nabcfd\n1223112:111\n"),
	}
	parameters := [][]string{
		{"-f", "1", "-s"},
		{"-f", "2,3", "-d", ":"},
	}

	for _, param := range parameters {
		for _, req := range requests {
			sysOut := executeReq(t, "cut", req, param)
			programOut := executeReq(t, "go", req, slices.Insert(param, 0, "run", "task.go"))
			if sysOut != programOut {
				t.Errorf(
					"answer wrong, system result:\n------\n%s------\nprogram result:%s\n",
					sysOut,
					programOut,
				)
			}
		}
	}

}
