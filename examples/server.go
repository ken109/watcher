package main

import (
	"bufio"
	"os"

	"github.com/ken109/watcher"
)

func main() {
	go watcher.Start(":9090")

	for {
		stdin := bufio.NewScanner(os.Stdin)
		if stdin.Scan() == false {
			return
		}

		err := watcher.Hook(stdin.Text())
		if err != nil {
			panic(err)
		}
	}
}
