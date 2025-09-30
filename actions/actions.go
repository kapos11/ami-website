package actions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

var AMIHost = "127.0.0.1:5038"

func SendAMIAction(action, command, username, secret string) (string, error) {
	conn, err := net.DialTimeout("tcp", AMIHost, 5*time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to connect to AMI: %v", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	var fullResponse strings.Builder

	loginCmd := fmt.Sprintf("Action: Login\r\nUsername: %s\r\nSecret: %s\r\n\r\n", username, secret)
	if _, err := writer.WriteString(loginCmd); err != nil {
		return "", err
	}
	writer.Flush()

	loginResp, _ := readAMIResponse(reader)
	fullResponse.WriteString(loginResp)

	var actionCmd string
	if command != "" {
		actionCmd = fmt.Sprintf("Action: %s\r\nCommand: %s\r\n\r\n", action, command)
	} else {
		actionCmd = fmt.Sprintf("Action: %s\r\n\r\n", action)
	}

	if _, err := writer.WriteString(actionCmd); err != nil {
		return fullResponse.String(), err
	}
	writer.Flush()

	// قراءة رد الأمر الرئيسي
	actionResp, err := readAMIResponse(reader)
	if err != nil {
		return fullResponse.String(), err
	}

	fullResponse.WriteString(actionResp)

	return fullResponse.String(), nil
}

func readAMIResponse(reader *bufio.Reader) (string, error) {
	var response strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return response.String(), err
		}
		response.WriteString(line)
		if strings.TrimSpace(line) == "" {
			break
		}
	}
	return response.String(), nil
}

func ParseResponse(amiResponse string) string {
	return amiResponse
}

func SendJSONAction(action, command, username, secret string) (string, error) {
	return SendAMIAction(action, command, username, secret)
}
