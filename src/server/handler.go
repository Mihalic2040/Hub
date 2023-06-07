package server

// Handler function type
type Handler func(input interface{}) (output interface{}, err error)

// HandlerMap holds a map of handler names to their corresponding functions
type HandlerMap map[string]Handler
