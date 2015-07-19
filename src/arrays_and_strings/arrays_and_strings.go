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
