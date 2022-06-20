package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	if client.Connect() != nil {
		log.Fatalln(client.Connect())
	}
	defer func() {
		if client.Connect() != nil {
			log.Fatalln(client.Connect())
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go worker(client.Receive, cancel)
	go worker(client.Send, cancel)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigCh:
		cancel()
		signal.Stop(sigCh)
		return

	case <-ctx.Done():
		close(sigCh)
		return
	}
}

func worker(handler func() error, cancel context.CancelFunc) {
	if err := handler(); err != nil {
		cancel()
	}
}
