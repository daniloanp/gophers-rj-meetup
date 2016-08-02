package app

import (
	baseApp "github.com/daniloanp/gophers-rj-meetup/singleton/app.v2"
)

type app struct {
	baseApp.App
	// additional fields ...
}

type App interface {
	baseApp.App
	Test() string
}

func (a *app) Test() string {
	return "Test @ App"
}

var appInstance *app = nil

func Instance() App {
	if appInstance == nil {
		appInstance = &app {
			App: baseApp.Instance(), // Ruim, estados dependem de outro módulo.
		}
	}
	return appInstance
}



