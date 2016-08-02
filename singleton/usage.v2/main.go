package main

import (
  "github.com/daniloanp/gophers-rj-meetup/singleton/app.v2"
  "fmt"
)



func main() {
  appInstance  := app.Instance()
  fmt.Println(appInstance.Calls(), &appInstance)
  // Agora já nao é possível instanciar app fora do módulo, no entanto, não podemos reutilizar de forma simplificada App.
}
