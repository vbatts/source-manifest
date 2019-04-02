package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "srcinfo",
	Short: "generate a bill-of-material from a filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		//...
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
	}
}
