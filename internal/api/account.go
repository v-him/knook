package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const accountEndpoint = lichessApi + "/account"
const emailEndpoint = accountEndpoint + "/email"

type Profile struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type Email struct {
	Email string `json:"email"`
}

func GetProfile(token string, client *http.Client) (Profile, error) {

	req, err := http.NewRequest(http.MethodGet, accountEndpoint, nil)
	if err != nil {
		return Profile{}, err
	}

	useToken(token, req)

	resp, err := client.Do(req)
	if err != nil {
		return Profile{}, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusTooManyRequests:
		waitMinute()
		return GetProfile(token, client)
	default:
		err := errors.New("I don't know how to handle status code: " + resp.Status)
		return Profile{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return Profile{}, err
	}

	profile := &Profile{}

	err = json.Unmarshal(body, profile)
	if err != nil {
		return Profile{}, err
	}

	return *profile, nil
}

func GetEmail(token string, client *http.Client) (string, error) {
	req, err := http.NewRequest(http.MethodGet, emailEndpoint, nil)
	if err != nil {
		return "", err
	}

	useToken(token, req)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusTooManyRequests:
		waitMinute()
		return GetEmail(token, client)
	default:
		err := errors.New("I don't know how to handle status code: " + resp.Status)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	emailObj := &Email{}

	err = json.Unmarshal(body, emailObj)
	if err != nil {
		return "", err
	}

	return emailObj.Email, nil
}
