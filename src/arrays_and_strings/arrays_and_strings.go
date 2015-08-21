package arrays_and_strings

//import "strings"
import "fmt"
import "sort"
import "strings"

//Write a function to detect if two strings are anagrams of each other.
func Detect_if_anagram(str1 string , str2 string) bool {

	var length1 = len(str1)
	var length2 = len(str2)

	if length1 != length2 {
				return false
	} else {
				str1 := Sort_String(str1)
				str2 := Sort_String(str2)

				fmt.Println("1. " , str1)
				fmt.Println("2. " , str2)

				if str1 == str2 {
					return true
				} else {
					return false
				}
			  return true
	}
}

//Write a function to detect if string has unique characters.
func Detect_If_String_Has_Unique_Chars(str string) bool {
  //1. itterate through all the chars and check to see if the char exist in the map if it does return false if not return true.
  //var myMap map[byte]bool
  myMap := make(map[byte]bool)

  //loop through string
  /*fmt.Println("Ivan was here", str)
  for i := 0 ; i < len(str) ; i++ {
    println(i)
    fmt.Printf("%c \n" ,str[i])
  }
  */


  for i := 0 ; i < len(str) ; i++ {
    //println(i)
    fmt.Printf("%c \n" ,str[i])
    if myMap[str[i]] == false  {
      myMap[str[i]] = true
    } else {
      return false
    }
  }
  return true
}

//Function to sort strings using slices?
func Sort_String(str string) string {
		newStr := strings.Split(str, "")
		sort.Strings(newStr)
		return strings.Join(newStr,"")
	}

//Reverse returns its argument string reversed rune-wise left to right.
func Reverse_String(s string) string {
  r := []rune(s)
  for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
  		r[i], r[j] = r[j], r[i]
  }
  return string(r)
}

//Fizz Buzz
/*
"Write a program that prints the numbers from 1 to 100. But for multiples of
three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
For numbers which are multiples of both three and five print “FizzBuzz”."
*/
func Fizz_Buzz(x int) {

	fmt.Println("*************************")
	fmt.Println("WELCOME TO THE FIZZ BUZZ!")
	fmt.Println("*************************")

	for i := 1 ; i < x ; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FIZZBUZZ")
		case i%5 == 0:
			fmt.Println("BUZZ")
		case i%3 ==0:
			fmt.Println("FIZZ")
		default:
			fmt.Println(i)
		}
	}
}

func Fizz_Buzz_2(x int) {
	fmt.Println("*************************")
	fmt.Println("WELCOME TO THE FIZZ BUZZ 2!")
	fmt.Println("*************************")

	for i := 1 ; i < x ; i++ {
		if i%15 == 0 {
			fmt.Println("FIZZ BUZZ")
		} else if i%5 == 0 {
			fmt.Println("BUZZ")
		} else if i%3 == 0 {
			fmt.Println("FIZZ")
		} else {
			fmt.Println(i)
		}
	}
}

// Write a function that returns the smallest number in a list.
func Find_Smallest(arr []int ) {
	fmt.Println("FIND SMALLEST VALUE")
	var smallest = 99999999
	for i:=1 ; i < len(arr) ; i++ {
			//fmt.Println(arr[i])
			if arr[i] <= smallest {
				smallest = arr[i]
			}
		}
		fmt.Println("Smallest: " , smallest )
}

// Write a function that returns the largest number in a list.
func Find_Largest(arr []int ) {
	fmt.Println("FIND LARGEST VALUE")
	var largest = -99999999
	for i:=1 ; i < len(arr) ; i++ {
			//fmt.Println(arr[i])
			if largest <= arr[i] {
				largest = arr[i]
			}
		}
		fmt.Println("Largest: " , largest )
}



/*

if x%15 == 0 {
	fmt.Println("FIZZ BUZZ")
} else if x%5 == 0 {
	fmt.Println("Shit")
} else if x%3 == 0 {
	fmt.Println("BUZZ")
} else {
	fmt.Println(i)
}


*/
