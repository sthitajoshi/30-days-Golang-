// 4-Read YAML config, run `docker ps`, log output
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	path := "example.txt"
	fileExists, _ := fileExists(path)
	fmt.Println("fileExist", fileExists)

	fmt.Println("file created", createFile(path))
	fmt.Println("file created if not exist:", createFileIfNot(path))
	fmt.Println("write files:", writeFile(path, []byte("hello buddy")))
	fmt.Println("append files:", appendFile(path, []byte("okkk that's it")))
}

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func createFile(path string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func createFileIfNot(path string) error {
	file, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_EXCL|os.O_CREATE,
		0666,
	)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func writeFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
func appendFile(path string, data []byte) error {
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

