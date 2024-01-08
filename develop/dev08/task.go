package main

import (
	"bufio"
	"fmt"
	psl "github.com/mitchellh/go-ps"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func cd(s string) {
	var err error
	if s == "" {
		err = os.Chdir(os.Getenv("HOME"))
	} else {
		err = os.Chdir(s)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

func kill(pidStr string) {
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = syscall.Kill(pid, syscall.SIGINT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Process %d killed", pid)
}

func ps() {
	procList, err := psl.Processes()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PID\tEXECUTABLE")
	for _, proc := range procList {
		info := proc.Executable()
		i := strings.Index(proc.Executable(), " ")
		if i == -1 {
			fmt.Printf("%d\t%s\n", proc.Pid(), info)
		} else {
			fmt.Printf("%d\t%s\n", proc.Pid(), info[:i])
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("enter command$ ")

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		line = line[:len(line)-1]
		if line == "" {
			continue
		}
		if line == "\\quit" {
			return
		}

		words := strings.Fields(line)
		switch words[0] {
		case "cd":
			if len(words) < 2 {
				cd("")
			} else if len(words) > 2 {
				fmt.Println("too many arguments")
			} else {
				cd(words[1])
			}
		case "pwd":
			pwd()
		case "echo":
			fmt.Println(strings.TrimPrefix(line, "echo"))
		case "kill":
			if len(words) < 2 {
				fmt.Println("enter pid")
			} else if len(words) > 2 {
				fmt.Println("too many arguments")
			} else {
				kill(words[1])
			}
		case "ps":
			if len(words) > 1 {
				fmt.Println("too many arguments")
			} else {
				ps()
			}
		default:
			cmd := exec.Command(words[0], words[1:]...)
			err := cmd.Run()
			cmd.Stdout = os.Stdout
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
