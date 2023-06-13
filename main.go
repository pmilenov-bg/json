package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func presentTree(entry, indent string, isLast bool) (error, []string, []string) {
	var listofFiles []string
	var listofFolders []string

	fileInfo, err := os.Stat(entry)
	if err != nil {
		return err, listofFolders, listofFolders
	}

	if fileInfo.IsDir() {
		fmt.Printf("%s├── %s\n", indent, fileInfo.Name())
		listofFolders = append(listofFolders, fileInfo.Name())

		fileInfoList, err := ioutil.ReadDir(entry)
		if err != nil {
			return err, listofFolders, listofFolders
		}

		fileCount := len(fileInfoList)
		for i, fileInfo := range fileInfoList {
			path := filepath.Join(entry, fileInfo.Name())

			var newIndent string
			var isLast bool

			if i == fileCount-1 {
				newIndent = indent + "    "
				isLast = true
			} else {
				newIndent = indent + "│   "
				isLast = false
			}
			err, childListofFolders, childListofFiles := presentTree(path, newIndent, isLast)
			listofFiles = append(listofFiles, childListofFiles...)
			listofFolders = append(listofFolders, childListofFolders...)

			if err != nil {
				return err, listofFolders, listofFolders
			}
		}
	} else {
		fmt.Printf("%s└── %s\n", indent, fileInfo.Name())
		listofFiles = append(listofFiles, fileInfo.Name())
	}

	return err, listofFolders, listofFiles //
}

func main() {
	// tempDir, err := ioutil.TempDir("", "example")
	// if err != nil {
	// 	fmt.Println("Failed to create temporary directory:", err)
	// 	return
	// }
	// defer os.RemoveAll(tempDir)
	// fmt.Println("Temporary directory created:", tempDir)
	directory := "/home/pete/-=Note=-"

	err, listofFolders, listofFiles := presentTree(directory, "", true)
	if err != nil {
		fmt.Printf("Error printing tree: %s\n", err.Error())
		return
	}
	fmt.Printf("%d folders, %d files\n", len(listofFolders), len(listofFiles))

}
