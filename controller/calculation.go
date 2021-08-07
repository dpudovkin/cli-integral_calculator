package calculation

import (
	"fmt"
	"math"
	"sync"

	"github.com/Knetic/govaluate"
)

func Calculate(a, b, delta float64, f func(float64) float64) float64 {
	var result float64
	for x := a; x <= b-delta; x += delta {
		result += delta * (f(x) + f(x+delta)) / 2
	}
	return result
}

func Concurrency_calculate(a, b, delta float64, f func(float64) float64) float64 {
	var result float64
	var cpu_num int = 4

	wg := &sync.WaitGroup{}
	wg.Add(cpu_num)
	channel := make(chan float64, cpu_num)

	step := math.Abs(b-a) / float64(cpu_num)

	for i := 0; i < cpu_num; i++ {
		go func(c chan float64, index int) {
			var start, end = a + float64(index)*step, a + float64(index+1)*step
			c <- Calculate(start, end, delta, f)
			wg.Done()
		}(channel, i)
	}

	wg.Wait()

	for i := 0; i < cpu_num; i++ {
		result += <-channel
	}

	return result
}

func MakeMathFunction(str string) func(float64) float64 {
	expression, err := govaluate.NewEvaluableExpression(str)
	if err != nil {
		fmt.Println("Invalid math function string")
		panic(err)
	}

	return func(x float64) float64 {
		parameters := make(map[string]interface{}, 1)
		parameters["x"] = x
		result, err := expression.Evaluate(parameters)
		if err != nil {
			return 0
		}
		convertValue, ok := result.(float64)
		if ok {
			return convertValue
		} else {
			return 0
		}
	}
}
