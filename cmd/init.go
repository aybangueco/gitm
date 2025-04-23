package cmd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize gitm",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing gitm...")

		if isInitialized() {
			fmt.Println("Gitm is already initialized!")
			return
		}

		err := initDB()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
