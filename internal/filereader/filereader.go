package filereader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	fmt.Println("Contents of file:", string(data))

	var lines = string(data)
	if len(lines) == 0 {
		return nil, errors.New("empty file")
	}

	return strings.Split(lines, "\n"), nil
}


func ReadFileBytes(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, err
	}
	//fmt.Println("Contents of file:", string(data))

	if len(data) == 0 {
		return nil, errors.New("empty file")
	}

	return data, nil
}