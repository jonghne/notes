package main

import (
	"fmt"
	"regexp"
)

func uncertain(in ...int) {
	fmt.Println(in)
}

func main() {
	valid := regexp.MustCompile(`^[1-9]\d{9}$`)

	fmt.Println(valid.MatchString("1234567890"))
	fmt.Println(valid.MatchString("0123456789"))
	fmt.Println(valid.MatchString("12345678900"))
	fmt.Println(valid.MatchString("1234567890a"))
	fmt.Println(valid.MatchString("123456789a"))
	fmt.Println(valid.MatchString("1234"))

	valid = regexp.MustCompile(`[\x00-\x60]`)
	fmt.Println(valid.MatchString("0xaF"))

	param := regexp.MustCompile(`:(.*)?`)
	fmt.Println(param.MatchString("/name/card"), param.FindStringSubmatch("/:name"))

	di := []int{2,3,4,5,6}
	uncertain(di...)
	uncertain(45,65)
}