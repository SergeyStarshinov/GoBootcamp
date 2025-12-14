package main

import (
	"context"
	"net/http"

	"tictactoe/internal/di"
	"tictactoe/internal/web/handler"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		di.Module,
		fx.Invoke(registerRoutes),
	)
	app.Run()
}

func registerRoutes(lc fx.Lifecycle, handler handler.GameHandler) {
	mux := createMux(handler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}

func createMux(handler handler.GameHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET 	/new_game", handler.CreateGameHumanHandler)
	mux.HandleFunc("GET 	/new_game_ai", handler.CreateGameAiHandler)
	mux.HandleFunc("POST 	/game/{x}", handler.NextMoveHandler)
	return mux
}
