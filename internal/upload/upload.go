package upload

import (
	"bufio"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
)

type Upload struct {
	File multipart.File
}

func (upload *Upload) ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (upload *Upload) ToText() ([]string, error) {
	id := uuid.NewV4()
	path := "./" + id.String() + ".txt"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, err
	}
	io.Copy(f, upload.File)
	dictionary, _ := upload.ReadLines(path)
	fmt.Printf("File has been saved in " + path)
	return dictionary, nil
}
