package main

import (
	"fmt"

	"github.com/ken109/watcher"
)

func handler(data interface{}) {
	fmt.Println(data)
}

func main() {
	go watcher.Watch("localhost:9090", handler)

	select {}
}
