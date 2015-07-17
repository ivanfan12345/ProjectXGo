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

//Function to sort strings using slices?
func Sort_String(str string) string {
		newStr := strings.Split(str, "")
		sort.Strings(newStr)
		return strings.Join(newStr,"")
	}
