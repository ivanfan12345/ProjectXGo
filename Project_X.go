// You can edit this code!
// Click here and start typing.
package main
import "fmt"
import "strings"


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
