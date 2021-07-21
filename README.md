# Example

![画面収録-2021-07-22-2 01 05](https://user-images.githubusercontent.com/37901879/126531762-9377a8f9-8474-46ae-8d0e-579f6e17667b.gif)


### Server

```go
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
```

### Client

```go
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
```
