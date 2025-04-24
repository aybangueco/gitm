package cmd

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrGitNotInstalled = errors.New("git not installed")
)
