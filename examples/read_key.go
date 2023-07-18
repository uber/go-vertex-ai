package examples

import (
	"fmt"
	"os"
)

// ReadKey is the function to read the secret json key from the file
func ReadKey() (string, error) {
	data, err := os.ReadFile(LicenseKeyFilePath)
	if err != nil {
		return "", fmt.Errorf("file reading error: %v", err)
	}

	return string(data), nil
}
