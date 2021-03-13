package question

import (
	question "github.com/bashmohandes/go-askme/question/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	question.NewRepository,
)
