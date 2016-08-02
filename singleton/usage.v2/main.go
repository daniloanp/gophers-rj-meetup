package main

import (
  "github.com/daniloanp/gophers-rj-meetup/singleton/app.v2"
  "fmt"
)



func main() {
  appInstance  := app.Instance()
  fmt.Println(appInstance.Dist(), &appInstance)
  // Agora já nao é possível instanciar app fora do módulo, no entanto, não podemos reutilizar de forma simplificada App.
  //appInstance2 := new (app.App)
  //fmt.Println((*appInstance2).Dist(), &appInstance2)
}
