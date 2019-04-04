package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// gendocCmd represents the gendoc command
var gendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, ".")
		if err != nil {
			glog.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(gendocCmd)
}
