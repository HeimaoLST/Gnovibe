package main

import (
	"context"
	"flag"
	"fmt"
	fswatch "fswatch/watch"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// path := flag.String()
	var pathFlag = flag.String("path", ".", "the path you want to watch")

	flag.Parse()
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	//need think
	sigchan := make(chan os.Signal, 1)
	//注册监听器
	signal.Notify(sigchan, syscall.SIGINT)
	wg.Add(1)
	go fswatch.Watch(ctx, pathFlag, &wg)

	<-sigchan
	fmt.Println("\nSee you next time!!")

	cancel()

	wg.Wait()

}
