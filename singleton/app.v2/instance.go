package app

type app struct {
  calls int64
  //....
}

type App interface {
  Calls() int64
}

func (a *app) Calls() int64 {
  a.calls += 1
  return a.calls
}

var appInstance *app = nil

func Instance() App {
  if appInstance == nil {
    appInstance = new(app)
  }
  return appInstance
}
