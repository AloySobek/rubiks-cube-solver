package algorithm

import (
	"bytes"
	"io"
	"os"
)

func WriteDataToFile(data *bytes.Buffer, filepath string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := file.Write(data.Bytes()); err != nil {
		panic(err)
	}
}

func ReadDataFromFile(filepath string) *bytes.Buffer {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer := bytes.NewBuffer([]byte{})

	io.Copy(buffer, file)

	return buffer
}
