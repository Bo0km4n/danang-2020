package cmd

import (
	"github.com/Bo0km4n/danang-2020/pkg"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ExecClient(srvHost, srvPort)
	},
}

var (
	srvPort string
	srvHost string
)

func init() {
	clientCmd.Flags().StringVarP(&srvPort, "port", "p", "8000", "remote host port")
	clientCmd.Flags().StringVarP(&srvHost, "host", "", "0.0.0.0", "remote host addr")
	root.AddCommand(clientCmd)
}
