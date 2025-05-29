package day4

import (
	"fmt"
	"io"
	"os"
)

func WriteToFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func AppendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, content)
	if err != nil {
		return err
	}

	return nil

}

func CopyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	/* content, err := io.ReadAll(file)
	if err != nil {
		return err
	} */

	destination, err := createFile(dst)
	if err != nil {
		return err
	}

	defer destination.Close()

	err = copyData(source, destination)
	if err != nil {
		return err
	}

	return nil

}

func createFile(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func copyData(src, dst *os.File) error {
	_, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func ListFiles(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	var content []string
	for _, entry := range entries {
		content = append(content, entry.Name())
	}

	return content, nil
}
