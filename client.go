package watcher

import (
	"encoding/json"
	"net"
	"time"
)

var clientConn net.Conn

func Watch(address string, fn func(data interface{})) {
	var err error
	connect(address)

	defer func() {
		err = clientConn.Close()
		if err != nil {
			return
		}
	}()

	buf := make([]byte, dataMaxLength)
	var jsonData interface{}

	for {
		_, err = clientConn.Write([]byte(""))
		if err != nil {
			panic(err)
		}

		n, err := clientConn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				connect(address)
				continue
			}
			panic(err)
		}

		err = json.Unmarshal(buf[:n], &jsonData)
		if err != nil {
			panic(err)
		}

		fn(jsonData)
	}
}

func connect(address string) {
	var err error

	for {
		clientConn, err = net.Dial("tcp", address)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
}
