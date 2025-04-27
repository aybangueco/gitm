package cmd

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var (
	colTitleID       = "ID"
	colTitleUsername = "Username"
	colTitleEmail    = "Email"
	colTitleActive   = "Active"
	rowHeader        = table.Row{colTitleID, colTitleUsername, colTitleEmail, colTitleActive}
)

var listAccountsCmd = &cobra.Command{
	Use:   "list-accounts",
	Short: "Display list of git accounts",
	Run: func(cmd *cobra.Command, args []string) {
		if !isInitialized() {
			cmd.Println("Gitm is not initialized yet, please run gitm init to initialize gitm!")
			return
		}

		t := table.NewWriter()
		t.AppendHeader(rowHeader)

		accounts, err := viewAccounts()
		if err != nil {
			cmd.Println(err.Error())
			return
		}

		for _, account := range accounts {
			t.AppendRow(table.Row{account.Id, *account.Username, *account.Email, *account.Active})
		}

		cmd.Println(t.Render())
	},
}

func init() {
	rootCmd.AddCommand(listAccountsCmd)
}
