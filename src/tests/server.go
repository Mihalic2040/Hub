package tests

func MyHandlerFromApp(input interface{}) (output interface{}, err error) {
	// Do some processing with the input data
	// ...

	// Return the output data and no error
	return input, nil
}

func MyHandlerFromAppp(input interface{}) (output interface{}, err error) {

	data := input.(int) * 2

	return data, nil
}

// func main() {
// 	//fake input
// 	input := app.InputData{
// 		HandlerName: "piska",
// 		Input:       50,
// 	}

// 	// runing server
// 	handlers := HandlerMap{
// 		app.GetFunctionName(MyHandlerFromApp): MyHandlerFromApp,
// 		"piska":                               MyHandlerFromAppp,
// 	}

// }
