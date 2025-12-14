package di

import (
	datamodel "tictactoe/internal/datasource/model"
	"tictactoe/internal/datasource/repository"
	dataservice "tictactoe/internal/datasource/service"
	gameservice "tictactoe/internal/domain/service"
	"tictactoe/internal/web/handler"
	webservice "tictactoe/internal/web/service"

	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(
		datamodel.NewStorage,
		repository.NewInMemoryRepository,
		dataservice.NewBaseDataService,
		gameservice.NewMinMax,
		webservice.NewBaseWebService,
		handler.NewBaseGameHandler,
	),
)
