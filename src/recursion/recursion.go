package recursion

import "fmt"


//Write a method to generate the nth Fibonacci Number recursively and itteratively.
func Get_Nth_Fib_Num_Recursive(nth int) int {
  if nth == 0 {
    return 0
    }
  if nth == 1 {
    return 1
    }
  return ( Get_Nth_Fib_Num_Recursive(nth - 1) + Get_Nth_Fib_Num_Recursive(nth -2) )
}

func Get_Nth_Fib_Num_Iter(nth int) int {
  if nth == 0 {
    return 0
    }
  if nth == 1 {
    return 1
    }

  var prev_prev = 0
  var prev = 1
  var result = 0

	for i:=2 ; i < nth ; i++ {
      result = prev + prev_prev
      prev_prev =  prev
      prev = result
  }
  return result
}


/*
The Fibonacci sequence is defined as: Fib(0) = 0, Fib(1) = 1, Fib(n) = Fib(n-1) + Fib(n-2).
Write a recursive function that can find Fib(n).
*/
func Print_Fibonacci(n int) int {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  } else {
    return ( Print_Fibonacci(n - 1) + Print_Fibonacci(n - 2))
  }
}

func Test() {
  fmt.Println("HELLO IVAN WAS HERE")
}
