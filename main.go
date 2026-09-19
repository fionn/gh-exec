package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// getToken is adapted from go-gh/pkg/auth/auth.go's tokenFromGh. This seems
// weird, the point of running as a gh extension is we inherit the
// authentication context -- but this is how go-gh derives the token (when using
// secure storage).
//
// The full flow for cli/gh is:
// * try to retrieve the token from the environment,
// * try to retrieve the token from disk,
// * try to retrieve the token from secure storage.
// These are all implemented on top of gh's configuration struct, so importing
// this functionality isn't ideal. So instead look at
// cli/go-gh/pkg/auth/auth.go, which we could use like
//
//	token, _ := auth.TokenForHost("github.com")
//
// to get the token, with an additional direct dependency. This would try to
// retrieve the token in the same order as gh, but on the last option, retrieval
// from secure storage (which we expect to be using in any case), it shells out
// to gh like we do (and wastes time checking if gh is executable, something we
// can safely skip).
func getToken() (string, error) {
	cmd := exec.Command("gh", "auth", "token")
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatal("Too few arguments")
	}

	token, err := getToken()
	if err != nil {
		log.Fatalf("Authentication token not found: %s", err.Error())
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
		log.Fatalf("Error executing command: %s", err.Error())
	}
}
