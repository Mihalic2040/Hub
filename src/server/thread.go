package server

import (
	"bufio"
	"log"

	"github.com/Mihalic2040/Hub/src/proto/api"
	"github.com/golang/protobuf/proto"
)

func Thread(handlers HandlerMap, rw *bufio.ReadWriter) {
	for {
		log.Println("Hello from thread")
		data, err := rw.ReadString('\n')
		if err != nil {
			break
		}

		data_raw := api.Response{}
		proto.Unmarshal([]byte(data), &data_raw)

		log.Println(data)
		log.Println(handlers)
	}

	// // Call a specific handler by name
	// handlerName := input.HandlerName
	// handler, ok := handlers[handlerName]
	// if !ok {
	// 	fmt.Printf("Handler '%s' not found\n", handlerName)
	// 	return
	// }

	// // Call the handler function with the input data
	// output, err := handler(inputData)
	// //handler(inputData)
	// if err != nil {
	// 	fmt.Printf("Error executing handler: %v\n", err)
	// 	return
	// }

	// // // Print the output data
	// fmt.Println(output)
}

// func Thread(handlers HandlerMap, rw *bufio.ReadWriter) {
// 	inputData := input.Input

// 	// Call a specific handler by name
// 	handlerName := input.HandlerName
// 	handler, ok := handlers[handlerName]
// 	if !ok {
// 		fmt.Printf("Handler '%s' not found\n", handlerName)
// 		return
// 	}

// 	// Call the handler function with the input data
// 	output, err := handler(inputData)
// 	//handler(inputData)
// 	if err != nil {
// 		fmt.Printf("Error executing handler: %v\n", err)
// 		return
// 	}

// 	// // Print the output data
// 	fmt.Println(output)
// }
