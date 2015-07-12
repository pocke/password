package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// TODO: path
	KEY_PATH = "/tmp/key"
)

func PasswordsFile() (*os.File, error) {
	// TODO: path
	path := "/tmp/passwords"
	return os.Open(path)
}

func Key(passphrase []byte) ([]byte, error) {
	f, err := os.Open(KEY_PATH)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	passkey := passphraseToAESKey(passphrase)
	c, err := NewCryptor(f, passkey)
	if err != nil {
		return nil, err
	}

	key, err := ioutil.ReadAll(c)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("key length should be 256bit. But got %d", len(key))
	}

	return key, nil
}

func NewKey(passphrase []byte) ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(KEY_PATH, os.O_WRONLY|os.O_CREATE, 0400)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	passKey := passphraseToAESKey(passphrase)
	c, err := NewCryptor(f, passKey)
	if err != nil {
		return nil, err
	}
	_, err = c.Write(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func passphraseToAESKey(passphrase []byte) []byte {
	res := make([]byte, 32)
	copy(res, passphrase)

	return res
}
