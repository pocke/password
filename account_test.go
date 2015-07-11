package main

import (
	"crypto/rand"
	"encoding/json"
	"reflect"
	"testing"
)

func TestLoadAccounts(t *testing.T) {
	m := newRWCMock()

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Error(err)
	}

	c, err := NewCryptor(m, key)
	if err != nil {
		t.Fatal(err)
	}

	expected := Accounts{
		{Name: "foo", Password: "piyopiyo", Icon: "aaa"},
		{Name: "nya", Password: "honya", Icon: "bbb"},
	}
	b, err := json.Marshal(expected)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Write(b)
	if err != nil {
		t.Fatal(err)
	}
	c.Close()

	a, err := LoadAccounts(m, key)
	if err != nil {
		t.Fatal(err)
	}

	for i, e := range expected {
		if !reflect.DeepEqual(e, a[i]) {
			t.Errorf("Expected: %+v, but got %+v", e, a[i])
		}
	}
}
