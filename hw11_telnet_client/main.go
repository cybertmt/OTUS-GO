package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var timeout time.Duration

const (
	minArgs        = 3
	defaultTimeout = 10
)

func init() {
	flag.DurationVar(&timeout, "timeout", defaultTimeout*time.Second, "connection timeout")
}

func main() {
	flag.Parse()

	if len(os.Args) < minArgs {
		log.Fatalf("Expected at least %d arguments", minArgs)
	}

	host, port := os.Args[2], os.Args[3]
	client := NewTelnetClient(net.JoinHostPort(host, port), timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		if err := client.Send(); err != nil {
			log.Printf("error during send: %v", err)
		} else {
			log.Printf("EOF")
		}
		cancel()
	}()

	go func() {
		if err := client.Receive(); err != nil {
			log.Printf("error during receive: %v", err)
		} else {
			log.Printf("connection was closed by peer")
		}
		cancel()
	}()

	<-ctx.Done()
}
