package parse

import (
	"errors"
	"flag"
	"fmt"
)

func Init() (cmd string, subCmd string, options any, parseError error) {
	showHelp := flag.Bool("help", false, "Show help message")
	showVersion := flag.Bool("version", false, "Show version")

	flag.Parse()

	if *showHelp {
		cmd = "help"
		return
	}

	if *showVersion {
		cmd = "version"
		return

	}

	cmd = flag.Arg(0)
	switch cmd {
	case "auth":
		authSubCmd, authOptions, err := auth()

		subCmd = authSubCmd
		options = authOptions
		parseError = err

	case "play":
	case "help":
	case "version":
	case "":
		parseError = errors.New("No command passed")

	default:
		errMsg := fmt.Sprintf("Unknown command: %s", cmd)
		parseError = errors.New(errMsg)
	}
	return
}
