package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func GetCurTime() (t time.Time, err error) {
	t, err = ntp.Time("0.beevik-ntp.pool.ntp.org")
	return
}

func main() {
	t, err := GetCurTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error occured while getting data: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Current time: %v\n", t)
}
