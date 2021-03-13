package answer

import (
	answer "github.com/bashmohandes/go-askme/answer/db"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	answer.NewRepository,
)
