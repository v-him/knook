package auth

import (
	"errors"
	"fmt"
	"os"
)

type LoginOptions struct {
	RequestToken bool
	Token string
}


func Login(opts LoginOptions) {
	token, found := os.LookupEnv("KNOOK_TOKEN")
	if !found || token == "" {
		err := errors.New("WARN: Variable KNOOK_TOKEN not found in the environment or was empty.")
		fmt.Println(err)
	}
}

