package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "hahah",
		Age:  18,
	}
	fmt.Println(u)
	up := &User{}
	fmt.Println(*up)
}
