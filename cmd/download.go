package cmd

import (
	"flow-cli/fdclient"
	"flow-cli/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)




var DownloadCmd = &cobra.Command{
    Use: "download",
    Long: "Download command, adds a new confirmed download",
    Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("this downloads something")
	fmt.Println(args)
	fdclient.NewDownloadConfirmed(args[0])
	if _, err := tea.NewProgram(tui.NewDownloadBar()).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}

    },
}
