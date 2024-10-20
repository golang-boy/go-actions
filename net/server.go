package net

import (
	"io"
	"net"
)

func Serve(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	for {
		bs := make([]byte, 8)
		_, err := conn.Read(bs)
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			return
		}

		// res := handleMsg(bs)
	}
}

func handleMsg(msg []byte) []byte {
	res := make([]byte, len(msg))

	copy(res[:len(msg)], msg)

	return res

}
