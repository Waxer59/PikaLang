package utils

import (
	"bufio"
	"os"
)

func ScanFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fileLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		fileLines = append(fileLines, line)
	}

	if err := scanner.Err(); err != nil {
		panic("There was an error reading the file: " + err.Error())
	}

	return fileLines, nil
}
