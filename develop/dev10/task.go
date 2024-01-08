package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type parsedFlags struct {
	timeout    time.Duration
	host, port string
}

var (
	protocol = "tcp"
	flags    = parsedFlags{}
)

func readSocket(conn net.Conn, done chan<- struct{}, errorCh chan<- error) {
	_, err := io.Copy(os.Stdout, conn)
	done <- struct{}{}
	if err != nil {
		errorCh <- err
	}
}

func writeSocket(conn net.Conn, done chan<- struct{}, errorCh chan<- error) {
	_, err := io.Copy(conn, os.Stdin)
	done <- struct{}{}
	if err != nil {
		errorCh <- err
	}
}

func main() {
	parseFlags()
	fmt.Println("Connecting...")
	conn, err := net.DialTimeout(protocol, flags.host+":"+flags.port, flags.timeout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	done := make(chan struct{})
	errorCh := make(chan error)

	go readSocket(conn, done, errorCh)
	go writeSocket(conn, done, errorCh)

	select {
	case err := <-errorCh:
		log.Fatalf("occured error:%v\n", err)
	case <-done:
		fmt.Println("Program stopped")
	}
}

func parseFlags() {
	flag.DurationVar(&flags.timeout, "timeout", 10*time.Second, "maximum time on create connection")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("not enough arguments, enter host and port")
	}
	flags.host = flag.Arg(0)
	flags.port = flag.Arg(1)
}
