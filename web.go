package main

import (
  "github.com/jmcarbo/gowut/gwu"
)

func main() {
  mw := gwu.NewWindow("main", "Main Windows")
  mw.Add(gwu.NewLabel("Hello World"))
  sr := gwu.NewServer("", ":9999")
  sr.AddWin(mw)
  sr.Start()
}
