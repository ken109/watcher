# Example

![watcher](https://user-images.githubusercontent.com/37901879/126532652-92ad5408-84c4-41cb-94c0-9bf38b2e6f06.gif)


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

func handler(data interface{}) {
	fmt.Println(data)
}

func main() {
	go watcher.Watch("localhost:9090", handler)

	select {}
}

```
