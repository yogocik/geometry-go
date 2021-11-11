package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	Scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Nama : ")
	Scanner.Scan()
	var name = Scanner.Text()

	fmt.Printf("Age : ")
	Scanner.Scan()
	var age = Scanner.Text()

	fmt.Println("Your name is", name)
	fmt.Println("Your age is", age)

	ageInt, _ := strconv.Atoi(age)

	fmt.Println("Age in number is", ageInt)
	fmt.Println("DONE with this scanner things! Hurraaaaaaaa.")
}
