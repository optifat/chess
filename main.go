package main

import (
    "log"

    "github.com/gotk3/gotk3/gtk"
)

func main() {
    gtk.Init(nil)

    b, err := gtk.BuilderNew()
    if err != nil {
        log.Fatal("Ошибка:", err)
    }

    err = b.AddFromFile("chess_gui.glade")
    if err != nil {
        log.Fatal("Ошибка:", err)
    }

    obj, err := b.GetObject("main_window")
    if err != nil {
        log.Fatal("Ошибка:", err)
    }

    win := obj.(*gtk.Window)
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    win.ShowAll()
    gtk.Main()
}
