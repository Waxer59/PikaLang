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

func MergeMaps[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
