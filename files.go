package main

import "os"

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir() //it is a file (not a folder)
}

func readTextFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeTextFile(path string, text string) error {
	return os.WriteFile(path, []byte(text), 0664)
}