package main

import (
	"fmt"
	"time"
)

func main() {
	s := "2020-02-02"
	t, _ := time.Parse(time.DateOnly, s)
	fmt.Println(t)
}
