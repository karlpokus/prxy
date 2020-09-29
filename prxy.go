package prxy

import (
	"errors"
	"io"
	"log"
	"net"
	"strings"
)

var argErr = errors.New("Argument error")

// list must contain the right things..
func validateArgs(list []string) (string, string, error) {
	var src, dest string
	if len(list) != 2 {
		return src, dest, argErr
	}
	if list[0] == "" || list[1] == "" {
		return src, dest, argErr
	}
	for _, s := range list {
		parts := strings.Split(s, ":")
		if len(parts) != 2 {
			return src, dest, argErr
		}
	}
	src = list[0]
	dest = list[1]
	return src, dest, nil
}

func Start(args []string) error {
	src, dest, err := validateArgs(args)
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", src)
	if err != nil {
		return err
	}
	log.Printf("listening on %s\n", src)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Accept connection err: %s", err)
			continue
		}
		log.Println("new client connection") // debug should be behind -v flag
		go handler(conn, dest)
	}
}

// handler calls dest and starts copying bytes between
// client and server
func handler(client net.Conn, dest string) {
	server, err := net.Dial("tcp", dest)
	if err != nil {
		log.Printf("Dial %s err: %s", dest, err)
		client.Close()
		return
	}
	log.Printf("new server connection to %s\n", dest)
	go copy(client, server, "client", "server")
	go copy(server, client, "server", "client")
}

// copy copies bytes from r to w and closes w if needed
func copy(w io.Writer, r io.Reader, writer, reader string) {
	io.Copy(w, r)
	log.Printf("%s ended connection\n", reader)
	// reader is closed in the other goroutine
	if v, ok := w.(io.WriteCloser); ok {
		v.Close()
		log.Printf("%s connection closed\n", writer)
	}
}
