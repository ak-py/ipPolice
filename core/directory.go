package core

import (
	"fmt"
	"os"
	"path/filepath"
)

const PathSeparator = "/"

func GetCurrentWorkingDirectory() string {
	wd, _ := os.Getwd()
	return wd
}

func getAssetsDirectory() string {
	cwd := ipPoliceConfig.WorkingDir
	path := filepath.FromSlash(cwd) + PathSeparator + "assets"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		parent := filepath.Dir(cwd)
		path = parent + PathSeparator + "assets"
	}
	return path
}

func dbFilePath() string {
	return filepath.FromSlash(getAssetsDirectory() + PathSeparator + DefaultDbFileName)
}

func logsDirectory() string {
	path := filepath.FromSlash(GetCurrentWorkingDirectory() + PathSeparator + "logs")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
		fmt.Println("Created logs directory")
	}

	return path
}
