package main

import (
	"os"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func gthread(f func()) {
	gdk.ThreadsEnter()
	defer gdk.ThreadsLeave()
	f()
}

func guiMain() {
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gdk.ThreadsEnter()
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("password")
	window.Connect("destroy", gtk.MainQuit)

	entry := gtk.NewEntry()

	vbox := gtk.NewVBox(false, 0)
	vbox.PackStart(entry, false, false, 1)

	window.Add(vbox)

	window.SetSizeRequest(460, 640)
	window.SetResizable(false)
	window.ShowAll()
	gtk.Main()
}
