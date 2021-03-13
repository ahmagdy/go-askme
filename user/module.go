package user

import (
	answer "github.com/bashmohandes/go-askme/answer/types"
	question "github.com/bashmohandes/go-askme/question/types"
	user "github.com/bashmohandes/go-askme/user/db"
	usecase "github.com/bashmohandes/go-askme/user/usecase"
	userRepo "github.com/bashmohandes/go-askme/user/types"
	"go.uber.org/fx"
)

var Module = fx.Provide(
		New,
		user.NewRepository,
)

type Params struct{
	fx.In

	userRepo userRepo.Repository
	questionRepo question.Repository
	answerRepo answer.Repository
}


type Result struct{
	fx.Out

	AuthUsecase    usecase.AuthUsecase
	AnswersUsecase usecase.AnswersUsecase
	AsksUsecase    usecase.AsksUsecase
}

func New(p Params)Result{
	return Result{
		AuthUsecase:    usecase.NewAuthUsecase(p.userRepo),
		AnswersUsecase: usecase.NewAnswersUsecase(p.questionRepo, p.answerRepo, p.userRepo),
		AsksUsecase:    usecase.NewAsksUsecase(p.questionRepo, p.answerRepo, p.userRepo),
	}
}
