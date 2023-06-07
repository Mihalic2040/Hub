package server

import (
	"fmt"

	"github.com/Mihalic2040/Hub/src/types"
)

func Thread(handlers HandlerMap, input types.InputData) {
	inputData := input.Input

	// Call a specific handler by name
	handlerName := input.HandlerName
	handler, ok := handlers[handlerName]
	if !ok {
		fmt.Printf("Handler '%s' not found\n", handlerName)
		return
	}

	// Call the handler function with the input data
	output, err := handler(inputData)
	//handler(inputData)
	if err != nil {
		fmt.Printf("Error executing handler: %v\n", err)
		return
	}

	// // Print the output data
	fmt.Println(output)
}
