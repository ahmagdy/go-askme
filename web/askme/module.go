package askme

import (
	"github.com/bashmohandes/go-askme/web/askme/controllers"
	"github.com/bashmohandes/go-askme/web/framework"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	New,
	controllers.Module,
	)

type Params struct{
	fx.In

	FrameworkApp framework.App
	HC *controllers.HomeController
	PC *controllers.ProfileController
	AC *controllers.OktaController


}

func New(p Params) *App{
	return NewApp(p.FrameworkApp, p.HC, p.PC, p.AC)
}