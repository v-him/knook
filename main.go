package main

import (
	"fmt"
	"os"

	"github.com/v-him/knook/internal/cmd/auth"
	"github.com/v-him/knook/internal/parse"
)

func main() {
	cmd, subCmd, opts, err := parse.Init()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse Error: %s", err.Error())
		os.Exit(1)
	}

	mainSwitch(cmd, subCmd, opts)
}

func mainSwitch(cmd string, subCmd string, opts any) {
	switch cmd {
	case "auth":
		switchAuth(subCmd, opts)
	case "play":
	case "help":
	case "version":
	}
}

func switchAuth(subCmd string, opts any) {
	switch subCmd {
	case "status":
		statusOpts, ok := opts.(auth.StatusOptions)
		if !ok {
			panic("Uncaught parsing error")
		}
		auth.Status(statusOpts)
	case "login":
		loginOpts, ok := opts.(auth.LoginOptions)
		if !ok {
			panic("Uncaught parsing error")
		}
		auth.Login(loginOpts)
	}
}
