package store

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func topicNameIsValid(topic string) error {
	if topic == "" {
		return errors.New("Empty topic name")
	}

	if l := len(strings.Split(topic, " ")); l == 0 || l > 1 {
		return errors.New("Invalid topic name")
	}

	return nil
}

// It returns the file name for the message file for a given topic
func getMessageFileName(topic string) string {
	return topic + ".msg"
}

// Custom openFile function to create the file in a predefined directory, if it doesn't exist else opening the existing file
func openFile(fileName string, flag int, perm fs.FileMode) (*os.File, error) {
	// Define the folder and file path
	dir := "files"
	filePath := filepath.Join(dir, fileName)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return nil, err
	}

	// Open the file, create it if it doesn't exist
	file, err := os.OpenFile(filePath, flag, perm)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return nil, err
	}

	return file, nil
}
