package main

import (
	"encoding/json"
	"io"
)

type Account struct {
	Name     string `json:"name"`
	Password string `json:"password"`

	// Icon is Base64 encoded image.
	Icon string `json:"icon"`
}

type Accounts []Account

func LoadAccounts(file io.ReadWriteCloser, key []byte) ([]Account, error) {
	c, err := NewCryptor(file, key)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	var res Accounts
	err = json.NewDecoder(c).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
