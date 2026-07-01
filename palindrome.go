package main
import "fmt"

func main(){
	var s string
	fmt.Println("Enter a string: ")
	fmt.Scan(&s)
	s2 := ""
	for i:= len(s)-1; i>=0; i--{
		s2 += string(s[i])
	}
	if s==s2{
		fmt.Println("palindrome")
	}else{
		fmt.Println("not palindrome")
	}
}