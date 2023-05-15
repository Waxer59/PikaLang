package utils

import (
	"io/ioutil"

	"github.com/inancgumus/screen"
)

func ScanFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func CallClearConsoleSc() {
	// Clears the screen
	screen.Clear()
	// Moves the cursor to the top left corner of the screen
	screen.MoveTopLeft()

}
