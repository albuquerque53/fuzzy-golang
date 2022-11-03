package app

import (
	"fmt"
	"fuzzygolang/internal/math"
	"fuzzygolang/internal/writer"
	"log"
)

type App struct {
	fNumber   int
	sNumber   int
	operation string
}

func StartApp() *App {
	operation := AskValue("Please, type the operation you want: ")
	validOp := validateOperation(operation)

	fNumber := AskValue("Type the first number of operation: ")
	sNumber := AskValue("The second number of operation: ")

	fmtFNumber, fmtSNumber := validateNumbers(fNumber, sNumber)

	return &App{
		fmtFNumber,
		fmtSNumber,
		validOp,
	}
}

func (app *App) Run() {
	var v int

	switch app.operation {
	case "+":
		v = math.Sum(app.fNumber, app.sNumber)
	case "-":
		v = math.Sub(app.fNumber, app.sNumber)
	case "*":
		v = math.Multiply(app.fNumber, app.sNumber)
	case "/":
		v = math.Div(app.fNumber, app.sNumber)
	default:
		log.Fatal("Invalid operation! The valid types are: +, -, * and / only")
	}

	res := writer.WriteJSON(app.fNumber, app.sNumber, app.operation, v)
	fmt.Println(res)
}
