package server

import (
	"fmt"
	"http-server-go/internal/parser"
	"log"
	"net"
)

// REMINDER: functions should only do one task.
func ListenForConnections() error {
	fmt.Println("Listening for connections on port 8080...")

	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		return err
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return err
	}

	msgFrmClient := make([]byte, 1024)
	n, err := conn.Read(msg_frm_sender)
	if err != nil {
		return err
	}
	defer conn.Close()

	parser.ParseRequest(msgFrmClient)
}
