package server

import (
	"fmt"
	"http-server-go/internal/http1"
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

	// TODO: Accept multiple connections
	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	msgFrmClient := make([]byte, 1024)
	n, err := conn.Read(msg_frm_sender)
	if err != nil {
		return err
	}

	// NOTE: should still return response when there are errors in parsing
	// NOTE: httpMessage here may have an empty body
	httpMessage, err := parser.ParseRequest(msgFrmClient)
	if err != nil {
		// check what kind of error returned and construct respective http response
	}

	httpResponse, err := http1.CreateResponse(httpMessage)
}
