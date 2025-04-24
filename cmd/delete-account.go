package cmd

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteAccountCmd = &cobra.Command{
	Use:   "delete-account [accountID]",
	Short: "Delete git account",
	Run: func(cmd *cobra.Command, args []string) {
		if !isInitialized() {
			cmd.Println("Gitm is not initialized yet, please run gitm init to initialize gitm!")
			return
		}

		id := args[0]

		parsedId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			cmd.Println("Invalid account ID (must be integer)")
			return
		}

		_, err = viewAccountByID(int(parsedId))
		if err != nil {
			if errors.Is(err, ErrAccountNotFound) {
				cmd.Println("Error: Account not found")
				return
			}
		}

		err = isGitInstalled()
		if err != nil {
			if errors.Is(err, ErrGitNotInstalled) {
				cmd.Println("Error: git not installed")
				return
			}

			cmd.Println("Error: ", err)
			return
		}

		err = deleteAccount(int(parsedId))
		if err != nil {
			cmd.Println(err.Error())
			return
		}

		cmd.Println("Account deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteAccountCmd)
}
