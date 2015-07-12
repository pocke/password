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

func guiMain(ch <-chan Accounts) {
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gdk.ThreadsEnter()
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("password")
	window.Connect("destroy", gtk.MainQuit)

	entry := gtk.NewEntry()

	store := gtk.NewListStore(glib.G_TYPE_STRING)
	treeview := gtk.NewTreeView()
	treeview.SetModel(store)
	// treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("icon", gtk.NewCellRendererPixbuf(), "pixbuf", 0))
	treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("name", gtk.NewCellRendererText(), "text", 0))

	go func() {
		for accounts := range ch {
			gthread(func() {
				store.Clear()
				for _, a := range accounts {
					var iter gtk.TreeIter
					store.Append(&iter)
					store.Set(&iter, 0, a.Name)
				}
			})
		}
	}()

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.Add(treeview)

	vbox := gtk.NewVBox(false, 0)
	vbox.PackStart(entry, false, false, 1)
	vbox.PackStart(swin, true, true, 0)

	window.Add(vbox)

	window.SetSizeRequest(460, 640)
	window.SetResizable(false)
	window.ShowAll()
}
