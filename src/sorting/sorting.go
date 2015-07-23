package sorting
//import "strings"
//import "fmt"
import "sort"
import "fmt"
//import "strings"
//Slow Sorts O(N^2)
/***** Selection Sort
Not stable
O(1) extra space
Θ(n2) comparisons
Θ(n) swaps
Not adaptive
******/
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
/*
Stable
O(1) extra space
O(n2) comparisons and swaps
Adaptive: O(n) time when nearly sorted
Very low overhead
*/
func Insertion_Sort(input []int) []int {
  return input
}


//Bubble Sort O(N^2)

// Optimal Sorts
// Merge Sort O(n log(n))
/*
Stable
Θ(n) extra space for arrays (as shown)
Θ(lg(n)) extra space for linked lists
Θ(n·lg(n)) time
Not adaptive
Does not require random access to data
*/
func Merge_Sort(data []int) []int {
  if len(data) == 1 {
    return data
  }

  //DIVIDE
  middle := len(data)/2
  //left of mid
  left := Merge_Sort(data[:middle])
  right := Merge_Sort(data[middle:])
  return Merge(left,right)
}

func Merge(left []int, right []int) []int {
  left_index := 0
  right_index := 0
  index := 0
  merged_array := make([]int, 0 , len(left) + len(right))

  fmt.Println(left_index)
  fmt.Println(right_index)
  fmt.Println(left)
  fmt.Println(right)
  fmt.Println(len(merged_array))


  for left_index < len(left) && right_index < len(right) {
    fmt.Println(index)
    if left[left_index] < right[right_index] {
      fmt.Println("HELLO IVAN")
      merged_array[index] = left[left_index]
      index++
      left_index++
    } else {
      merged_array[index] = right[right_index]
      index++
      right_index++
    }
  }




  for left_index < len(left) {
    merged_array[index] = left[left_index]
    index++
    left_index++
  }
  for right_index < len(right) {
    merged_array[index] = right[right_index]
    index++
    right_index++
  }
  return merged_array
}


// Quick Sort O(n log(n))
