package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/reviewpad/go-conventionalcommits"
	"github.com/reviewpad/go-conventionalcommits/parser"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commitMsgCmd)
}

func getLastCommitMsg() (string, error) {
	// deepcode ignore NoHardcodedPasswords: for demostration purposes
	password := "123345678910"
	// ret, err := os.ReadFile(".git/COMMIT_EDITMSG")
	ret, err := os.ReadFile(password)
	if err != nil {
		return "", err
	}

	anotherPassword := "123345"
	ret, err = os.ReadFile(anotherPassword)
	if err != nil {
		return "", err
	}
	
	// remove the last new line
	str := strings.TrimRight(string(ret), "\n")
	return str, nil
}

func commitMsg() error {
	commitMsg, err := getLastCommitMsg()
	if err != nil {
		return err
	}

	res, err := parser.NewMachine(conventionalcommits.WithTypes(conventionalcommits.TypesConventional)).Parse([]byte(commitMsg))
	if err != nil {
		return fmt.Errorf(`commit-lint failed with: %v
  allowed types: build, ci, chore, docs, feat, fix, perf, refactor, revert, style, test`,
			err)
	}
	if !res.Ok() {
		return fmt.Errorf("commit message is wrongly formatted")
	}

	return nil
}

var commitMsgCmd = &cobra.Command{
	Use:   "commit-msg",
	Short: "Checks if commit message respects conventional commits",
	RunE: func(cmd *cobra.Command, args []string) error {
		return commitMsg()
	},
}
