package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func initializeConfigVariables() {
	HOME_DIR, _ := os.UserHomeDir()
	CONFIG_DIR = filepath.Join(HOME_DIR, ".config/"+APP_NAME)
	CONFIG_FILE = filepath.Join(CONFIG_DIR, "config.json")
}

func configFolderExists() bool {
	_, err := os.Stat(CONFIG_DIR)
	return !os.IsNotExist(err)
}

func createConfigFolder() {
	os.Mkdir(CONFIG_DIR, 0755)
}

func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
