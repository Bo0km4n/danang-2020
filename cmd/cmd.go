package cmd

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "danang",
	Short: "da",
}

func Execute() error {
	return root.Execute()
}
