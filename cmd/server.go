package cmd

import (
	"github.com/Bo0km4n/danang-2020/pkg"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "srv",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.ExecServer(listenPort)
	},
}

var (
	listenPort string
)

func init() {
	serverCmd.Flags().StringVarP(&listenPort, "port", "p", "8000", "listen port")
	root.AddCommand(serverCmd)
}
