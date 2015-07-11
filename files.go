package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	KEY_PATH = "/tmp/key"
)

func PasswordsFile() (*os.File, error) {
	// TODO: path
	path := "/tmp/passwords"
	return os.Open(path)
}

func Key() ([]byte, error) {
	res, err := ioutil.ReadFile(KEY_PATH)
	if err != nil {
		return nil, err
	}
	if len(res) != 32 {
		return nil, fmt.Errorf("key length should be 256bit. But got %d", len(res))
	}
	return res, nil
}

func NewKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(KEY_PATH, key, 0400)
	if err != nil {
		return nil, err
	}

	return key, nil
}
