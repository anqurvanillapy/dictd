package main

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"github.com/anqurvanillapy/dictd/dictd"
)

// Serve the database
func Serve(port int, conf ClusterConf) error {
	addr := fmt.Sprintf(":%d", port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("Listening to", addr, "...")

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	var res dictd.Response
	msg := bytes.TrimRight(buf, "\x00\t\r\n ")
	cmd := dictd.ParseMsg(string(msg))

	if cmd == nil {
		res = dictd.StatusFail
		log.Printf("Invalid request %q\n", msg)
	} else {
		res = cmd.Execute()
		log.Printf("Request: %q, Response: %q\n", cmd, res)
	}

	conn.Write([]byte(res))
	conn.Close()
}
