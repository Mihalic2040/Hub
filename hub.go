package hub

import (
	"context"

	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
)

type App struct {
	types.App
}

// AddHandler adds a handler to the Handlers map in the Host struct
func (app *App) Handler(name string, handlerFunc server.Handler) {
	if app.Handlers == nil {
		app.Handlers = make(server.HandlerMap)
	}
	app.Handlers[name] = handlerFunc
}

// Conf sets the configuration in the App struct
func (app *App) Settings(conf types.Config) {
	app.Config = conf
}

func (app *App) Start(serve bool) *types.App {
	ctx := context.Background()
	host := node.Start_host(ctx, app.Config, app.Handlers, serve)
	app.App = host
	return &host
}
