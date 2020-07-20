package main

import (
    "log"
    "strconv"
    "C"
    "github.com/gotk3/gotk3/gtk"
)

var buffer int
var fieldButtons [64]*gtk.Button
var fieldPieces [64]*gtk.Image

func buttonClicked(button *gtk.Button, index int){
    if buffer == -1{
        buffer = index
    }else if buffer != index{
        fieldPieces[index].SetFromPixbuf(fieldPieces[buffer].GetPixbuf())
        fieldPieces[buffer].Clear()
        buffer = -1
    }
}

func main() {
    gtk.Init(nil)

    b, err := gtk.BuilderNew()
    if err != nil {
        log.Fatal("Error:", err)
    }

    err = b.AddFromFile("chess_gui.glade")
    if err != nil {
        log.Fatal("Error:", err)
    }

    obj, err := b.GetObject("main_window")
    if err != nil {
        log.Fatal("Error:", err)
    }

    win := obj.(*gtk.Window)
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    obj, _ = b.GetObject("white")
    whiteInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("black")
    blackInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("queens")
    queensInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("rooks")
    rooksInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("bishops")
    bishopsInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("knights")
    knightsInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("pawns")
    pawnsInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("kings")
    kingsInput := obj.(*gtk.Entry)

    obj, _ = b.GetObject("submit_button")
    submitButton := obj.(*gtk.Button)

    buffer = -1

    for i := 0; i <= 63; i++{
        obj, _ = b.GetObject("button_"+strconv.Itoa(i))
        fieldButtons[i] = obj.(*gtk.Button)
        obj, _ = b.GetObject(strconv.Itoa(i))
        fieldPieces[i] = obj.(*gtk.Image)
    }


    for index, button := range fieldButtons{
        button.Connect("clicked", buttonClicked, index)
    }

    var bitboards Bitboard

    submitButton.Connect("clicked", func() {
        whiteStr, _ := whiteInput.GetText()
        bitboards.white, _ = strconv.ParseUint(whiteStr, 10, 64)
        blackStr, _ := blackInput.GetText()
        bitboards.black, _ = strconv.ParseUint(blackStr, 10, 64)
        queensStr, _ := queensInput.GetText()
        bitboards.queens, _ = strconv.ParseUint(queensStr, 10, 64)
        rooksStr, _ := rooksInput.GetText()
        bitboards.rooks, _ = strconv.ParseUint(rooksStr, 10, 64)
        bishopsStr, _ := bishopsInput.GetText()
        bitboards.bishops, _ = strconv.ParseUint(bishopsStr, 10, 64)
        knightsStr, _ := knightsInput.GetText()
        bitboards.knights, _ = strconv.ParseUint(knightsStr, 10, 64)
        pawnsStr, _ := pawnsInput.GetText()
        bitboards.pawns, _ = strconv.ParseUint(pawnsStr, 10, 64)
        kingsStr, _ := kingsInput.GetText()
        bitboards.kings, _ = strconv.ParseUint(kingsStr, 10, 64)

        for i := 0; i <= 63; i++{
            if (bitboards.white >> i)%2 == 1{
                if (bitboards.queens >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/white_queen.png")
                }else if (bitboards.rooks >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/white_rook.png")
                }else if (bitboards.bishops >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/white_bishop.png")
                }else if (bitboards.knights >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/white_knight.png")
                }else if (bitboards.pawns >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/white_pawn.png")
                }else if (bitboards.kings >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/white_king.png")
                }
            }else if (bitboards.black >> i)%2 == 1{
                if (bitboards.queens >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/black_queen.png")
                }else if (bitboards.rooks >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/black_rook.png")
                }else if (bitboards.bishops >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/black_bishop.png")
                }else if (bitboards.knights >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/black_knight.png")
                }else if (bitboards.pawns >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/black_pawn.png")
                }else if (bitboards.kings >> i)%2 == 1{
                    fieldPieces[i].SetFromFile("images/black_king.png")
                }
            }
        }
	})

    win.ShowAll()
    gtk.Main()
}
