package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// getToken is adapted from pkg/auth/auth.go's tokenFromGh. This seems weird,
// the point of running as a gh extension is we inherit the authentication
// context -- but this is literally how the gh extension itself derives the
// token.
func getToken() (string, error) {
	cmd := exec.Command("gh", "auth", "token")
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}

func main() {
	if len(os.Args) < 2 {
		panic("Too few arguments")
	}

	token, err := getToken()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Env = append(cmd.Environ(), fmt.Sprintf("GITHUB_TOKEN=%s", token))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
		panic(err)
	}
}
