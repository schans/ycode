package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
)

// Expects ykman on the path
func getAccounts() ([]string, error) {
	out, err := exec.Command("ykman", "oath", "accounts", "list").Output()
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Cmd out:\n%s\n", out)
	return strings.Split(strings.ReplaceAll(string(out), "\r\n", "\n"), "\n"), nil
}

// Expects ykman on the path
func getCode(account string) (string, error) {
	out, err := exec.Command("ykman", "oath", "accounts", "code", "--single", account).Output()
	if err != nil {
		return "", err
	}
	//fmt.Printf("Cmd out:\n%s\n", out)
	return strings.TrimSuffix(string(out), "\n"), nil
}

func main() {
	// get yubico accounts
	accounts, err := getAccounts()
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}
	//fmt.Printf("Accounts: %v\n", accounts)

	// setup answer
	answers := struct {
		Account string
	}{}

	// setup question
	accountQs := []*survey.Question{
		{
			Name: "account",
			Prompt: &survey.Select{
				Message: "Choose an account:",
				Options: accounts,
			},
			Validate: survey.Required,
		},
	}

	// ask the question
	err = survey.Ask(accountQs, &answers)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}

  // get code
	code, err := getCode(answers.Account)
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}
	clipboard.WriteAll(code)
	//fmt.Printf("the code is %s\n", code)

	// print the answers
	fmt.Printf("Copied code to clipboard for  %s.\n", answers.Account)
}
