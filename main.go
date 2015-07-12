package main

import (
	"time"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	accounts := make(chan Accounts)

	go func() {
		accounts <- Accounts{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "baz"},
		}
		time.Sleep(1 * time.Second)
		accounts <- Accounts{
			{Name: "hoge"},
			{Name: "fuga"},
			{Name: "piyo"},
		}
	}()

	guiMain(accounts)
	if !KeyExist() {

	}
	gtk.Main()
}
