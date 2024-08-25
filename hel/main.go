package main

import (
	"fmt"
	"github.com/longsizhuo/forum/models"
)

func main() {
	const a int = 1
	fmt.Printf("a = %d\n", a)
	fmt.Println("Hello, World!")
	print("Hello, World!")
	fmt.Println(add(1, 2))

	c := add(1, 2)
	fmt.Println(c)
	d := models.Admin{AdminID: a, AdminName: "admin", AdminPassword: "admin", RoleID: 1, Status: 1}
	fmt.Println(d)

}

func add(a int, b int) int {
	return a + b
}
