package main

import (
"fmt"
"regexp"
)

func main(){
	str := "11golang11"
	match,_ := regexp.MatchString("g([a-z]+)g",str)
	fmt.Println(match)

	str = "user.12345"
	match,_ = regexp.MatchString("user.([0-9]{6,20})",str)
	fmt.Println(match)
}
