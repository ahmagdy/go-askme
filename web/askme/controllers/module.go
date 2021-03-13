package controllers

import (
	user "github.com/bashmohandes/go-askme/user/usecase"
	"github.com/bashmohandes/go-askme/web/framework"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Params struct{
	fx.In

	Router framework.Router
	Renderer framework.Renderer
	Config *framework.Config
	SessionManager framework.SessionManager
	AuthUsecase user.AuthUsecase
	AsksUsecase user.AsksUsecase
	AnswersUsecase user.AnswersUsecase
}

type Result struct{
	fx.Out

	HomeController *HomeController
	OktaController *OktaController
	ProfileController *ProfileController
}

func New(p Params) Result{
	return Result{
		HomeController: NewHomeController(p.Router),
		OktaController: NewOktaController(p.Router, p.Renderer, p.Config, p.SessionManager, p.AuthUsecase),
		ProfileController: NewProfileController(p.Router, p.Renderer, p.AsksUsecase, p.AnswersUsecase),
	}
}