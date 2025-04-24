package cmd

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
)

func isInitialized() bool {
	if _, err := os.Stat("gitm.db"); err == nil {
		return true
	}

	return false
}

func isGitInstalled() error {
	_, err := exec.LookPath("git")
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return ErrGitNotInstalled
		}

		return err
	}

	return nil
}

func validateUsername(s string) error {
	if len(s) < 4 {
		return errors.New("Username too short")
	}

	if len(s) > 30 {
		return errors.New("Username too long")
	}
	return nil
}

func validateEmail(s string) error {
	emailPattern := `^(?i)([a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}|[0-9]+[+][a-z0-9-]+@users\.noreply\.github\.com)$`

	isMatch, err := regexp.MatchString(emailPattern, s)
	if err != nil {
		return err
	}

	if !isMatch {
		return errors.New("Invalid email format")
	}

	return nil
}
