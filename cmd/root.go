package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Slack-Noti-Job",
		Short: "Slack-Noti-Job",
		Long:  "Slack-Noti-Job",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("hello world")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
