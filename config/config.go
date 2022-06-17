package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// To communicate with the sia renter we need to have it's API password.
// By default this is located at ~/.sia/apipassword and this builds the
// absolute path
func GetAPIPasswordPath(path string) string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(currentUser.HomeDir, path)
}

func PrependDataDirectory(path string) string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome != "" {
		return filepath.Join(dataHome, "sia-nbdserver", path)
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(currentUser.HomeDir, ".local/share/sia-nbdserver", path)
}

// Get unix domain socket to connect to
func GetSocketPath(path string) string {
	runtimeDir := os.Getenv("XDG_RUNTIME_DIR")
	if runtimeDir == "" {
		log.Fatal("$XDG_RUNTIME_DIR not set")
	}
	return filepath.Join(runtimeDir, path)
}

func ReadPasswordFile(path string) (string, error) {
	passwordBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", nil
	}

	return strings.TrimSpace(string(passwordBytes)), nil
}
