// You can edit this code!
// Click here and start typing.
package main
import "fmt"
import "strings"
import "sort"

func main() {
	fmt.Println(" ========== Welcome to ProjectX Go ==========\n\n")
	Detect_Hacker_Name()

	//TEST
	var str1 = "ivan"
	var str2 = "navi"
	//str := []string{"ivan","navi"}

	fmt.Println(Detect_if_anagram(str1 , str2))
}

func Detect_Hacker_Name() {
	fmt.Print("Enter your alias: ")
	var input string
	fmt.Scanf("%s",&input)
	fmt.Println("Welcome ", input ,"let's hack!")
}

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
