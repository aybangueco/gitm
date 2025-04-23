package cmd

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var (
	username string
	email    string
)

var addAccountCmd = &cobra.Command{
	Use:   "add-account",
	Short: "Add new git account",
	Run: func(cmd *cobra.Command, args []string) {
		if !isInitialized() {
			fmt.Println("Gitm is not initialized yet, please run gitm init to initialize gitm!")
			return
		}

		addAccountForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Title("What is your git username?").Value(&username).Validate(validateUsername).Placeholder("myawesomeusername"),
				huh.NewInput().Title("What is your git email?").Value(&email).Validate(validateEmail).Placeholder("myemailiscool@gmail.com"),
			),
		)

		err := addAccountForm.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = addNewAccount(account{Username: username, Email: email})
		if err != nil {
			fmt.Println("Error creating new account: ", err.Error())
			return
		}

		fmt.Println("Account created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addAccountCmd)
}
