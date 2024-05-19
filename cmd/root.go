package cmd

import (
	"flow-cli/fdclient"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


func init(){
    rootCmd.AddCommand(DownloadCmd)
}


var rootCmd = &cobra.Command{
    Use: "flow",
    Short: "A fast download manager powered by rust",
    Long: "A blazing fast download manager using a rust daemon and a go cli",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to flow")
	fdclient.SubscribeToDownloads()
    },
}


func Execute(){
 if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
