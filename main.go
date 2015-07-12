package main

import "time"

func main() {
	ch := make(chan Accounts)
	go func() {
		ch <- Accounts{
			{Name: "foo"},
			{Name: "bar"},
			{Name: "baz"},
		}
		time.Sleep(1 * time.Second)
		ch <- Accounts{
			{Name: "hoge"},
			{Name: "fuga"},
			{Name: "piyo"},
		}
	}()
	guiMain(ch)
}
