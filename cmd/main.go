package main

import (
	"math/rand"
	"time"

	"github.com/BigJk/snd/printing/cups"
	"github.com/BigJk/snd/printing/remote"
	"github.com/BigJk/snd/printing/serial"

	"github.com/BigJk/snd/server"
)

var serverOptions []server.Option
var startFunc = startServer

func startServer() {
	rand.Seed(time.Now().UnixNano())

	s, err := server.NewServer("./data.db", append(serverOptions, server.WithPrinter(&cups.CUPS{}), server.WithPrinter(&remote.Remote{}), server.WithPrinter(&serial.Serial{}))...)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	_ = s.Start(":7123")
}

func main() {
	startFunc()
}
