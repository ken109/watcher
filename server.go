package watcher

import (
	"encoding/json"
	"errors"
	"net"
)

const dataMaxLength = 4 * 1024

var listener net.Listener
var serverConns []net.Conn

func Start(port string) {
	var err error
	listener, err = net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		serverConns = append(serverConns, conn)
	}
}

func Hook(data interface{}) error {
	if len(serverConns) > 0 {
		jsonData, err := json.Marshal(data)

		if len(jsonData) > dataMaxLength {
			return errors.New("data is too large")
		}

		for _, conn := range serverConns {
			_, err = conn.Write(jsonData)
		}

		return err
	}
	return nil
}
