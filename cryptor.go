package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

type Cryptor struct {
	stream  cipher.Stream
	storage io.ReadWriteCloser
}

func NewCryptor(rwc io.ReadWriteCloser, key []byte) (*Cryptor, error) {
	c := &Cryptor{storage: rwc}

	err := c.setStream(key)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Cryptor) setStream(key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	var iv [aes.BlockSize]byte
	c.stream = cipher.NewOFB(block, iv[:])
	return nil
}

func (c *Cryptor) Read(b []byte) (int, error) {
	return cipher.StreamReader{S: c.stream, R: c.storage}.Read(b)
}

func (c *Cryptor) Write(b []byte) (int, error) {
	return cipher.StreamWriter{S: c.stream, W: c.storage}.Write(b)
}

func (c *Cryptor) Close() error {
	return c.storage.Close()
}

var _ io.ReadWriteCloser = &Cryptor{}
