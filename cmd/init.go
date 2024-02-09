package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var hooks = []string{"commit-msg"}

func getCurrentGitRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	err = isGitRepository(cwd)
	if err != nil {
		return "", err
	}

	return cwd, nil
}

func isGitRepository(path string) error {
	f, err := os.Lstat(filepath.Join(path, ".git"))
	if err != nil {
		return err
	}
	if !f.IsDir() {
		return fmt.Errorf(".git must be a directory")
	}
	return nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Install git hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		gitRoot, err := getCurrentGitRoot()
		if err != nil {
			return err
		}

		for _, githook := range hooks {
			script := fmt.Sprintf("#!/bin/sh\n%v %v\n", rootCmd.Use, githook)
			hookFile := path.Join(gitRoot, ".git/hooks", githook)
			if err := os.WriteFile(hookFile, []byte(script), os.ModePerm); err != nil {
				return err
			}
			fmt.Printf("Installing %v\n", hookFile)
		}

		return nil
	},
}
