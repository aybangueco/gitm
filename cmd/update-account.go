package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var updateAccountCmd = &cobra.Command{
	Use:   "update-account [accountID]",
	Short: "Update existing git account",
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

		account, err := viewAccountByID(int(parsedId))
		if err != nil {
			if errors.Is(err, ErrAccountNotFound) {
				cmd.Println("Error: Account not found")
				return
			}
		}

		updateAccountForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("What is your git username?").Value(*&account.Username).Validate(validateUsername).Placeholder("yournewawesomeusername"),
				huh.NewInput().Title("What is your git email?").Value(*&account.Email).Validate(validateEmail).Placeholder("mynewcoolemail@gmail.com"),
			),
		)

		err = updateAccountForm.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = updateAccount(int(parsedId), account)
		if err != nil {
			cmd.Println(err.Error())
			return
		}

		cmd.Println("Account updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateAccountCmd)
}
