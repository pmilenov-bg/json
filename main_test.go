package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var tempDir string

func setupTest() {
	tempDir, _ := ioutil.TempDir("", "example")

	fmt.Print(tempDir)

	fmt.Printf("%s ...the temp dir name", tempDir)

	createDirectory(tempDir)

}

func tearDown() {
	log.Println("teardown test")
	fmt.Println("Teardown Stage")
	os.RemoveAll(tempDir)
}

func createDirectory(tempDir string) (string, error) {

	fmt.Println("Temporary directory created:", tempDir)

	// Create test directory structure
	err := os.Mkdir(filepath.Join(tempDir, "pete"), 0755)
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

	err = ioutil.WriteFile(filepath.Join(tempDir, "pete", "cv.txt"), []byte("this_is_cv_file"), 0644)
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

func TestPresentTree(t *testing.T) {
	fmt.Println("Setup Stage")

	_, files, folders := presentTree(tempDir, "", true)

	if len(files) != 9 {
		fmt.Printf("%d folders, %d files\n", len(files), len(folders))
		t.Error("Ops")
	}

}

func TestMain(m *testing.M) {
	setupTest()
	code := m.Run()

	os.Exit(code)
}
