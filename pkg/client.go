package pkg

import (
	"fmt"
	"net"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/sirupsen/logrus"
)

func ExecClient(host string, port string) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	logrus.Infof("Connected to %s", conn.RemoteAddr().String())
	run(conn)
	return nil
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "quit", Description: "Close this connection"},
	}
	return s
}

func executor(in string) {
	if in == "quit" {
		logrus.Info("Close connection")
		os.Exit(0)
	}
	if _, err := cli.conn.Write([]byte(in)); err != nil {
		logrus.Error(err)
	}
}

var cli *client

func run(conn net.Conn) {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(conn.RemoteAddr().String()+" >>> "),
		prompt.OptionTitle("danang-tcp-client"),
	)
	cli = &client{conn: conn, p: p}
	cli.p.Run()
}

type client struct {
	conn net.Conn
	p    *prompt.Prompt
}
