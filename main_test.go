package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func setup() {
	fmt.Println("Setup Stage")
	// Perform any setup operations here
}

func teardown() {
	fmt.Println("Teardown Stage")
	// Perform any teardown operations here
}

func createTestDirectoryStructure() (string, error) {
	tempDir, err := ioutil.TempDir("", "example")
	if err != nil {
		return "", err
	}

	// Create test directory structure
	err = os.Mkdir(filepath.Join(tempDir, "pete"), 0755)
	if err != nil {
		return "", err
	}
	err = os.Mkdir(filepath.Join(tempDir, "pete/surf"), 0755)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filepath.Join(tempDir, "pete/surf", "general.txt"), []byte("general_file_name"), 0644)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(filepath.Join(tempDir, "pete", "cv.txt"), []byte("this_is_cv_file_name"), 0644)
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

// func test
// func testMain(m *testing.M) {
// 	setup()
// 	defer teardown()

// 	m.Run()
// }
