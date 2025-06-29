package parse

import (
	"errors"
	"flag"
	"fmt"
	"os"

	authCmd "github.com/v-him/knook/internal/cmd/auth"
)

func auth() (string, any, error) {
	authSubCmd := flag.Arg(1)
	switch authSubCmd {
	case "status":
		opts := authStatus()
		return authSubCmd, opts, nil
	case "":
		err := errors.New("No subcommand passed to auth")
		return "", nil, err
	default:
		errMsg := fmt.Sprintf("Unknown auth subcommand: %s", authSubCmd)
		return "", nil, errors.New(errMsg)
	}
}

func authStatus() authCmd.StatusOptions {
	authStatusCmd := flag.NewFlagSet("auth status", flag.ExitOnError)
	quiet := authStatusCmd.Bool("quiet", false, "If command prints output or just sets exit code")
	authStatusCmd.Parse(os.Args[3:])
	return authCmd.StatusOptions{Quiet: *quiet}
}

