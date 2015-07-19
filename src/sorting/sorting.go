package sorting

//import "strings"
//import "fmt"
//import "sort"
//import "strings"

//Slow Sorts O(N^2)
func Selection_Sort(input []int) []int {
  /*for i := 0 ; i < len(input) ; i++ {
    fmt.Println(input[i])
  }
  */

  last := len(input) - 1
  for i := 0 ; i < last ; i ++ {
    input_min :=  input[i]
    i_min := i

    for j := i + 1 ; j < len(input) ; j++ {
      if input[j] < input_min {
        input_min = input[j]
        i_min = j
      }
    }
    input[i] , input[i_min] = input_min , input[i]
  }
  return input
}
