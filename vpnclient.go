package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var content *fyne.Container = container.New(layout.NewVBoxLayout())

func index() {
    signinbutton := widget.NewButton("Sign In", signinbutton)
    loginbutton := widget.NewButton("Log In", loginbutton)
    content.Add(layout.NewSpacer())
    content.Add(signinbutton)
    content.Add(loginbutton)
    content.Add(layout.NewSpacer())
}


func signinbutton() {
    content.RemoveAll()
}

func loginbutton() {
    content.RemoveAll()
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

    index()

	w.SetContent(content)

    w.ShowAndRun()
    tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
