package sorting
//import "strings"
//import "fmt"
import "sort"
import "fmt"
//import "strings"
//Slow Sorts O(N^2)
//Selection Sort
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
    fmt.Println(input)
  }
  return input
}

func Selection_Sort_Interface(array sort.Interface) {
  last := array.Len() - 1
  for i := 0 ; i < last ; i++ {
    i_min := i
    for j := i + 1 ; j < array.Len(); j++ {
      if array.Less(j , i_min) {
        i_min = j
      }
    }
    array.Swap(i,i_min)
    fmt.Println("SWAP!" , array)
  }
}
//Insertion Sort O(N^2)

//Bubble Sort O(N^2)

// Optimal Sorts
// Merge Sort O(n log(n))

// Quick Sort O(n log(n))
