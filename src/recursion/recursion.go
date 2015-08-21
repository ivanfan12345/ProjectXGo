package recursion

import "fmt"

/*
The Fibonacci sequence is defined as: Fib(0) = 0, Fib(1) = 1, Fib(n) = Fib(n-1) + Fib(n-2).
Write a recursive function that can find Fib(n).
*/
func Fib(n int) int {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  } else {
    return ( Fib(n - 1) + Fib(n - 2))
  }
}

func Test() {
  fmt.Println("HELLO IVAN WAS HERE")
}
