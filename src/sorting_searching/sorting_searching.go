package sorting_searching
//import "strings"
//import "fmt"
import "sort"
import "fmt"
//import "strings"

/************************* Selection Sort *************************
Not stable
O(1) extra space
Θ(n2) comparisons
Θ(n) swaps
Not adaptive
************************* Selection Sort *************************/
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
/************************* Insertion Sort *************************
O(N^2)
/*
Stable
O(1) extra space
O(n2) comparisons and swaps
Adaptive: O(n) time when nearly sorted
Very low overhead
************************** Insertion Sort *************************/
func Insertion_Sort(input []int) []int {
  return input
}

/************************* Bubble Sort *************************
 O(N^2)
// Optimal Sorts
// Merge Sort O(n log(n))
/*
Stable
Θ(n) extra space for arrays (as shown)
Θ(lg(n)) extra space for linked lists
Θ(n·lg(n)) time
Not adaptive
Does not require random access to data
************************* Bubble Sort *************************/
func Merge_Sort(data []int) []int {
  if len(data) == 1 {
    return data
  }

  //DIVIDE
  middle := len(data)/2
  //left of mid
  left := Merge_Sort(data[:middle])
  right := Merge_Sort(data[middle:])
  data = Merge(left,right)
  fmt.Println(data)
  return data
}

func Merge(left []int, right []int) []int {
  left_index := 0
  right_index := 0
  index := 0
  merged_array := make([]int, 0)

  for left_index < len(left) && right_index < len(right) {
    fmt.Println(index)
    if left[left_index] < right[right_index] {
      fmt.Println("HELLO IVAN")
      //[index] = left[left_index]
      merged_array = append(merged_array,left[left_index])
      index++
      left_index++
    } else {
      //merged_array[index] = right[right_index]
      merged_array = append(merged_array,right[right_index])
      index++
      right_index++
    }
  }
  for left_index < len(left) {
    //merged_array[index] = left[left_index]
    merged_array = append(merged_array,left[left_index])
    index++
    left_index++
  }
  for right_index < len(right) {
    //merged_array[index] = right[right_index]
    merged_array = append(merged_array,right[right_index])
    index++
    right_index++
  }
  return merged_array
}

// Quick Sort O(n log(n))

//Find two pairs that sum up to target.
func Find_Two_Sum(data []int , target int ) []int {
  var difference = 0
  pair := make([]int, 2)
  var m = map[int]int{}
  for i , element := range data {
    fmt.Println("i:",i)
    fmt.Println("element:", element)
    if m[element] != 0 {
      //found pair
      pair[0] = m[element]
      pair[1] = element
      fmt.Println("FOUND PAIR!!!")
      return pair
    } else {
      //store difference
      difference = target - element
      m[element] = difference
    }
  }
  fmt.Println("NOT FOUND!!! :(")
  return pair
}

//Binary Search Iter and Recursively
func Binary_Search_Recur(data []int , target int , low int , high int) int {
  if high < low {
    return -1
  }
  mid := (low + high)/2
  if data[mid] > target {
    return Binary_Search_Recur(data , target , low , mid-1)
  } else if data[mid] < target {
    return Binary_Search_Recur(data , target , mid + 1 , high)
  }
  return mid
}
func Binary_Search_Iter(data []int, target int) int {
  low := 0
  high := len(data) -1
  for low <= high {
    mid := (high + low)/2
    if data[mid] > target {
      high =  mid - 1
    } else if data[mid] < target {
      low = mid + 1
    } else {
      return mid
    }
  }
  return -1
}
// You are given two sorted arrays, A and B, where A has a large enough buffer at
// the end to hold B. Write a method to merge B into A in sorted order.
func Merge_Two_Sorted_arrays(A []int , B []int) []int {
  merged_array := make([]int, len(A) +len(B))
  indexA := 0
  indexB := 0
  index := 0

  for indexA < len(A) && indexB < len(B) {
      if A[indexA] < B[indexB] {
        merged_array = append(merged_array,A[indexA])
        indexA++
        index++
      } else {
        merged_array =  append(merged_array,B[indexB])
        indexB++
        index++
      }
  }
  for indexA < len(A) {
    merged_array = append(merged_array,A[indexA])
    indexA++
    index++
  }
  for indexB < len(B) {
    merged_array = append(merged_array,B[indexB])
    indexB++
    index++
  }
  return merged_array
}





/***** CTCI
- Write a method to sort an array of strings so that all the anagrams are next to each other. _._ pg 361
- Given a sorted array of n integers that has been rotated an unknown number of times, write code to find an element in the array. You may assume that the array was originally sorted in increasing order.
       EXAMPLE
       Input: find 5 in {15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
       Output: 8 (the index of 5 in the array)
- Imagine you have a 20 GB file with one string per line. Explain how you would sort the file. pg 364
- Given a sorted array of strings which is interspersed with empty strings, write a method to find the location of a given string.
       EXAMPLE
       Input: find "ball" in {"at", "", "", "", "ball", "", "", "car", tti) tat ftf4^.j>j t(» <t»\ , aau , , )
       Output: 4

- Given an M x N matrix in which each row and each column is sorted in ascending order, write a method to find an element. pg 36S
- A circus is designing a tower routine consisting of people standing atop one another's shoulders. For practical and aesthetic reasons, each person must be both shorter and lighter than the person below him or her. Given the heights and weights of each person in the circus, write a method to compute the largest possible number of people in such a tower.
       EXAMPLE:
       Input (ht,wt): (65, 100) (70, 150) (56, 90) (75, 190) (60, 95) (68, 110)
       Output:The longest tower is length 6 and includes from top to bottom:
       (56, 90) (60,95) (65,100) (68,110) (70,150) (75,190)
       p j 1 ? :
 - Imagine you are reading in a stream of integers. Periodically, you wish to be able to look up the rank of a number x (the number of values less than or equal tox). Implement the data structures and algorithms to support these operations.That is, implement the method track(int x), which is called when each number is generated, and the method getRankOfNumber(int x), which returns the number of values less than or equal to x (not including x itself).
       EXAMPLE
       Stream (in order of appearance): 5, 1, 4, 4, 5, 9, 7, 13, 3 getRankOfNumber(l) = 0 getRankOfNumber(3) = 1 getRankOfNumber(4) = 3

******/
