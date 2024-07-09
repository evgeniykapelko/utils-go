package utils

import (
	"cli_cobra/logger"
	"fmt"
	"golang.org/x/term"
	"os"
)

func GetPassphrase() ([]byte, error) {
	fmt.Print("Enter passphrase: ")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("error reading passphrase: %v", err)
	}

	defer safeRestore(int(os.Stdin.Fd()), oldState)

	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("error reading passphrase: %v", err)
	}

	return passphrase, nil
}

func safeRestore(fd int, oldState *term.State) {
	if err := term.Restore(fd, oldState); err != nil {
		logger.HaltOnErr(fmt.Errorf("error restoring terminal state: %v", err))
	}
}
