package metadata

import (
	"os"
	"strings"
)

func readFile(name string, fallback string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		return fallback
	}
	return strings.TrimSpace(string(data))
}

func Version() string {
	return readFile("VERSION", "development")
}

func Build() string {
	return readFile("BUILD", "unknown")
}

func Codename() string {
	return readFile("CODENAME", "development")
}
