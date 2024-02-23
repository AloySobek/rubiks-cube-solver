package solver

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"
)

func writeDataToFile(data *bytes.Buffer, filepath string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := file.Write(data.Bytes()); err != nil {
		panic(err)
	}
}

func readDataFromFile(filepath string) *bytes.Buffer {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	buffer := bytes.NewBuffer([]byte{})

	io.Copy(buffer, file)

	return buffer
}

func mapToBytes(m map[uint64]int) *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(m); err != nil {
		panic(err)
	}

	return buffer
}

func bytesToMap(b *bytes.Buffer) map[uint64]int {
	var m map[uint64]int

	decoder := gob.NewDecoder(b)

	if err := decoder.Decode(&m); err != nil {
		panic(err)
	}

	return m
}
