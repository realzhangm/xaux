package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	proxyAddr = flag.String("aa", "127.0.0.1:11024", "asr ai_proxy address")
)

func init() {
	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	sc, err := NewSoundCap(ctx, *proxyAddr)
	if err != nil {
		panic(err)
	}
	go func() {
		err = sc.Run()
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGILL, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-sigChan

	cancel()
	fmt.Println("signal = ", sig)
	sc.Release()
	sc.Dump("xx.pcm")
}
