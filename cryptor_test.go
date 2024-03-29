package main

import (
	"bytes"
	"crypto/rand"
	"io/ioutil"
	"reflect"
	"testing"
)

type rwcMock struct {
	*bytes.Buffer
	Closed bool
}

func (m *rwcMock) Close() error {
	m.Closed = true
	return nil
}

func newRWCMock() *rwcMock {
	return &rwcMock{Buffer: bytes.NewBuffer([]byte{})}
}

func TestCryptor(t *testing.T) {
	m := newRWCMock()
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Error(err)
	}

	c, err := NewCryptor(m, key)
	if err != nil {
		t.Error(err)
	}
	c.Write([]byte("hoge"))
	c.Close()

	buf := bytes.NewBuffer(m.Buffer.Bytes())
	m = &rwcMock{Buffer: buf}
	c, err = NewCryptor(m, key)
	if err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadAll(c)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(b, []byte("hoge")) {
		t.Errorf("Expected: %q, but got %q", []byte("hoge"), b)
	}
}
