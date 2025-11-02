package fswatch

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watch(ctx context.Context, path *string, wg *sync.WaitGroup) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	watcher.Add(*path)
	defer watcher.Close()
	log.Println("Start")
	for {
		select {
		case event, ok := <-watcher.Events:
			{
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					fmt.Printf("[CREATE] %s at %s\n", event.Name, time.Now().Format("2006-01-02 15:04:05"))
				} else if event.Has(fsnotify.Remove) {
					fmt.Printf("[DELETE] %s at %s\n", event.Name, time.Now().Format("2006-01-02 15:04:05"))

				} else if event.Has(fsnotify.Write) {
					fmt.Printf("[MODIGY] %s at %s\n", event.Name, time.Now().Format("2006-01-02 15:04:05"))

				}

			}
		case err, ok := <-watcher.Errors:
			{
				if !ok {
					return
				}

				log.Println("error:", err)

			}
		case <-ctx.Done():
			{
				wg.Done()
				return
			}
		}

	}
}
