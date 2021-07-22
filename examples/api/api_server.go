package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ken109/watcher"
)

func hookHandler(c *gin.Context) {
	err := watcher.Hook(c.Query("key"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func main() {
	go watcher.Start(":9090")

	r := gin.Default()

	r.GET("", hookHandler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
