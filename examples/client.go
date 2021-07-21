package main

import (
	"fmt"

	"github.com/ken109/watcher"
)

func main() {
	go watcher.Watch(
		":9090",
		func(data interface{}) {
			fmt.Println(data)
		},
	)

	select {}
}
