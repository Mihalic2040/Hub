package server

import "github.com/Mihalic2040/Hub/src/proto/api"

// Handler function type
type Handler func(input *api.Request) (response api.Response, err error)

// HandlerMap holds a map of handler names to their corresponding functions
type HandlerMap map[string]Handler
