package cmd

import (
	"errors"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var switchAccountCmd = &cobra.Command{
	Use:   "switch-account [accountID]",
	Short: "Switch git account",
	Args:  cobra.ExactArgs(1),
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

		acc, err := viewAccountByID(int(parsedId))
		if err != nil {
			if errors.Is(err, ErrAccountNotFound) {
				cmd.Println("Error: Account not found")
				return
			}
		}

		activeAccount, err := getActiveAccount()
		if err != nil {
			cmd.Println(err.Error())
			return
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

		gitUsernameCmd := exec.Command("git", "config", "--global", "user.name", *acc.Username)
		if err := gitUsernameCmd.Run(); err != nil {
			cmd.Println("Failed to set git username:", err)
			return
		}

		gitEmailCmd := exec.Command("git", "config", "--global", "user.email", *acc.Email)
		if err := gitEmailCmd.Run(); err != nil {
			cmd.Println("Failed to set git email:", err)
			return
		}

		oldAccountActive := false
		err = updateAccount(activeAccount.Id, account{Active: &oldAccountActive})
		if err != nil {
			cmd.Println(err.Error())
			return
		}

		switchedAccountActive := true
		err = updateAccount(int(parsedId), account{Active: &switchedAccountActive})
		if err != nil {
			cmd.Println(err.Error())
			return
		}

		cmd.Println("Account switched successfully!")
	},
}

func init() {
	rootCmd.AddCommand(switchAccountCmd)
}
