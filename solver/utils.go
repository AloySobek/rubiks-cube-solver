package solver

import (
	"bytes"
	"encoding/gob"
	"io"
	"os"
)

func writeDataToFile(data *bytes.Buffer, filepath string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write(data.Bytes()); err != nil {
		return err
	}

	return nil
}

func readDataFromFile(filepath string) (*bytes.Buffer, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	buffer := bytes.NewBuffer([]byte{})

	io.Copy(buffer, file)

	return buffer, nil
}

func mapToBytes(m map[uint64]int) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer([]byte{})

	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(m); err != nil {
		return nil, err
	}

	return buffer, nil
}

func bytesToMap(b *bytes.Buffer) (map[uint64]int, error) {
	var m map[uint64]int

	decoder := gob.NewDecoder(b)

	if err := decoder.Decode(&m); err != nil {
		return nil, err
	}

	return m, nil
}

func fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}

	return true
}
