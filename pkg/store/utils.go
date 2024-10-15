package store

import (
	"errors"
	"strings"
)

func topicNameIsValid(topic string) error {
	if topic == "" {
		return errors.New("Empty topic name")
	}

	if l := len(strings.Split(topic, "")); l == 0 || l > 1 {
		return errors.New("Invalid topic name")
	}

	return nil
}

// It returns the file name for the message file for a given topic
func getMessageFileName(topic string) string {
	return "files/" + topic + ".msg"
}
