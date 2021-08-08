# Golang CLI appllication for calculating integral
## Description
This application allows you to find a first-order integral for any function on a given interval using the trapezoid method.  
The user sets the boundaries of the interval, the expression for calculating the function with the argument x, and the step width, 
which directly affects the accuracy of the result obtained.
## Examples 
```console
$ go run main.go integrate 0 10 0.001 x*x/2
Result:  166.62604874949946
```
```console
$ go run main.go integrate 1 6 0.00001 x/2+1/x+2*x*x
Result:  153.8747105016243
```
## Concurrency test
To achieve better performance, the value of the integral is calculated in parallel. 
The initial interval is divided into equal intervals, then the program calculates the value of integrals for these intervals and summarizes the resulting values.
The number of threads is equal to the number of processor cores.
```console
$ go run main.go 
1 6 0.000001 1-x+1/x
-10.708235698204676
No-concurrency - 4435 millisecond 
-10.708229943952627
With concurrency - 2381 millisecond
```
## Additional info
```console
$ go run main.go integrate --help

Calculation of the integral of a user-defined function over a certain interval using the trapezoidal method.

Usage:
   main <a> <b> <delta> <function> {flags}

Arguments: 
   a                             Initial value of the interval
   b                             End value of the interval
   delta                         Step size. The less the more accurate
   function                      Integrated function. For example x*x, x/2, sqrt(log(x)) etc.

Flags: 
   -h, --help                    displays usage information of the application or a command (default: false)
```
