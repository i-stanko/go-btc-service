package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	subscribersFile = "subscribers.txt"
)

func isEmailExists(email string) (bool, error) {
	file, err := os.Open(subscribersFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == email {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}

func saveEmail(email string) error {
	file, err := os.OpenFile(subscribersFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(email + "\n")
	if err != nil {
		return err
	}

	return writer.Flush()
}
