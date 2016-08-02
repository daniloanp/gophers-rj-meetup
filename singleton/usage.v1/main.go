package main

import (
  "github.com/daniloanp/gophers-rj-meetup/singleton/app.v1"
  "fmt"
)



func main() {
  appInstance  := app.Instance()
  fmt.Println(appInstance.Calls(), &appInstance)

  // Versão problemática porque nada me impede de criar novas instâncias
  appInstance2 := new (app.App)
  fmt.Println(appInstance2.Calls(), &appInstance2)
}
