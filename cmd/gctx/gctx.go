package main

import (
	"github.com/tomromeo/gctx/pkg/lib"

	"github.com/rivo/tview"
)

func main() {

    app := tview.NewApplication()
    list := tview.NewList().
        SetWrapAround(true)
    frame := tview.NewFrame(list)
    frame.SetBorder(true)
    frame.SetRect(0,0,50,20)
    frame.SetTitle("Select GitHub user")
    lib.AddVimNav(list)

    users := lib.ParseUserFile()

    for _, user := range users {
        list.AddItem(user.Username, user.Email, 0, nil)
    }
    list.SetSelectedFunc(func(i int, _ string, _ string, _ rune)  {
        if (i != -1) {
            lib.ApplyUser(users[i])
        }
        app.Stop()
    })
    if err := app.SetRoot(frame, false).SetFocus(list).Run(); err != nil {
        panic(err)
    }
}
