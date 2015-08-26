// You can edit this code!
// Click here and start typing.
package main
import "fmt"
import "recursion"
import "arrays_and_strings"
				//"sort"
				//"sorting"

func main() {
	fmt.Println(" ========== Welcome to ProjectX Go ==========\n\n")
	Detect_Hacker_Name()

	//TEST DATA
	//var str1 = "Ivan"
	//var str2 = "navI"
	//str := []string{"Ivan","navI"}
	//test := []int{-1,22,0,4,-4,21,1,111,1,-17,}
	//fmt.Println(arrays_and_strings.Detect_if_anagram(str1 , str2))
	//fmt.Println(arrays_and_strings.Detect_If_String_Has_Unique_Chars("abcdefghijklmnaop"))

	//fmt.Println("BEFORE: ",test)
	//test = sorting.Merge_Sort(sort.IntSlice(test));
	//fmt.Println("AFTER: ",test)
	//arrays_and_strings.Fizz_Buzz(99)
	//arrays_and_strings.Fizz_Buzz_2(100)

	//arr  := []int{66,55,4,-3,3,3,33,444}
	//arrays_and_strings.Find_Largest(arr)

	//var nth = 99
	//for i := 0 ; i < nth ; i++ {
	//	fmt.Println("fib(",i,")",recursion.Fib(i))
	//}

	fmt.Println(recursion.Get_Nth_Fib_Num_Iter(1))





}

func Detect_Hacker_Name() {
	fmt.Print("ENTER YOUR ALIAS: ")
	var input string
	fmt.Scanf("%s",&input)
	fmt.Println("Welcome ", input ,"let's hack!")
	fmt.Println("Your reversed hacker name is: ", arrays_and_strings.Reverse_String(input))
}
