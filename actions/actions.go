package actions

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var BaseURL = "http://127.0.0.1:8088/asterisk/manager"

var client = &http.Client{
	Timeout: 30 * time.Second,
}

// SendAction sends a request to Asterisk Web AMI
func SendAction(data url.Values) (string, error) {
	req, err := http.NewRequest("POST", BaseURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Asterisk: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}

// SendJSONAction converts JSON data to form values and sends it
func SendJSONAction(action, command, username, secret string) (string, error) {
	data := url.Values{}
	data.Set("action", action)
	if command != "" {
		data.Set("command", command)
	}
	data.Set("username", username)
	data.Set("secret", secret)

	resp, err := SendAction(data)
	if err != nil {
		return "", err
	}
	return resp, nil
}
