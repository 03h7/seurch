package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	files := []string{}
	directories := []string{}

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	defaultShortcutsDir := user.HomeDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs"

	dirEntries, err := os.ReadDir(defaultShortcutsDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			if strings.ToLower(entry.Name()) != "desktop.ini" {
				files = append(files, entry.Name())
			}
		} else {
			directories = append(directories, entry.Name())
		}
	}

	for _, directory := range directories {
		cwd := defaultShortcutsDir + "\\" + directory
		subDirEntries, err := os.ReadDir(cwd)
		if err != nil {
			fmt.Println(err)
		}
		files = drilldownDirectories(subDirEntries, files, cwd)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

func drilldownDirectories(subDirEntries []os.DirEntry, files []string, cwd string) []string {
	for _, subDirEntry := range subDirEntries {
		if !subDirEntry.IsDir() {
			if strings.ToLower(subDirEntry.Name()) != "desktop.ini" {
				files = append(files, subDirEntry.Name())
			}
		} else {
			cwd := cwd + "\\" + subDirEntry.Name()
			subDirEntries, err := os.ReadDir(cwd)
			if err != nil {
				fmt.Println(err)
			}
			files = drilldownDirectories(subDirEntries, files, cwd)
		}
	}
	return files
}
