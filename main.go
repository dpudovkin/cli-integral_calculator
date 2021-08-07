package main

import (
	calculation "concurrency_calculation/integral/controller"
	"fmt"
	"strconv"
	"time"

	"github.com/thatisuday/commando"
)

type IntegrateCommand struct {
}

func main() {
	configureCommando()
	commando.Parse(nil)
}

func consoleMode() {

	var a, b, delta float64
	var str string
	fmt.Scan(&a, &b, &delta, &str)
	f := calculation.MakeMathFunction(str)
	fmt.Println(calculation.Concurrency_calculate(a, b, delta, f))
}

func configureCommando() {

	commando.
		SetExecutableName("main").
		SetVersion("0.1").
		SetDescription("This tool concurrency calculate integral of any function")

	commando.
		Register("integrate").
		SetShortDescription("Calculating an integral over a certain interval").
		SetDescription("Calculation of the integral of a user-defined function over a certain interval using the trapezoidal method.").
		AddArgument("a", "Initial value of the interval", "").
		AddArgument("b", "End value of the interval", "").
		AddArgument("delta", "Step size. The less the more accurate", "").
		AddArgument("function", "Integrated function. For example x*x, x/2, sqrt(log(x)) etc.", "").
		SetAction(integrateAction)

}

func integrateAction(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
	f := calculation.MakeMathFunction(args["function"].Value)
	var a, b, delta float64
	var err error
	a, err = strconv.ParseFloat(args["a"].Value, 64)
	if err != nil {
		panic(err)
	}
	b, err = strconv.ParseFloat(args["b"].Value, 64)
	delta, err = strconv.ParseFloat(args["delta"].Value, 64)
	fmt.Println("Result: ", calculation.Concurrency_calculate(a, b, delta, f))
}

func concurrency_test(a, b, delta float64, f func(float64) float64) {
	timestamp := timeMillis()
	fmt.Println(calculation.Calculate(a, b, delta, f))
	fmt.Printf("Millisecond: %d\n", timeMillis()-timestamp)

	timestamp = timeMillis()
	fmt.Println(calculation.Concurrency_calculate(a, b, delta, f))
	fmt.Printf("Millisecond: %d\n", timeMillis()-timestamp)
}

func timeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
