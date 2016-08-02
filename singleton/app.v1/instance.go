package app

type App struct {
    calls int64
}

func (a *App) Calls() int64 {
  a.calls += 1
  return a.calls
}

var app *App = nil

func Instance() *App {
  if app == nil {
    app = new(App)
  }
  return app
}
