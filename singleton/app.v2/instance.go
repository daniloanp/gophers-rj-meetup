package app

type app struct {
  //....
}

type App interface {
  Dist() string
}

func (a *app) Dist() string {
  return "App @ stone"
}

var appInstance *app = nil

func Instance() App {
  if appInstance == nil {
    appInstance = new(app)
  }
  return appInstance
}
