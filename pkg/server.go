package pkg

import (
	"bufio"
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

func ExecServer(port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	logrus.Infof("Listening on %s", listen.Addr().String())
	for {
		conn, err := listen.Accept()
		logrus.Infof("Accepted connection from %s", conn.RemoteAddr().String())
		if err != nil {
			return err
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) error {
	defer conn.Close()
	conn.Write([]byte("HELLO\n"))
	reader := bufio.NewReader(conn)
	remoteAddr := conn.RemoteAddr().String()
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				logrus.Infof("Closed connection by %s", remoteAddr)
				break
			}
			return err
		}
		content := string(buf[:n])
		if isQuit(content) {
			logrus.Infof("Closed connection by %s", remoteAddr)
			return nil
		}
		logrus.Infof("Received content: %s", content)
	}
	return nil
}

func isQuit(line string) bool {
	switch line {
	case "QUIT":
		return true
	case "quit":
		return true
	}
	return false
}
