package framework

import "go.uber.org/fx"

var Module = fx.Provide(
	NewConnection,
	NewInMemorySessionStore,
	NewRenderer,
	NewRouter,
	NewApp,
	)
