package app

type App struct {
    //....
}

func (a *App) Dist() string {
  return "App @ stone"
}

var app *App = nil

func Instance() *App {
  if app == nil {
    app = new(App)
  }
  return app
}
