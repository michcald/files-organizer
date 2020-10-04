package main

import (
	"fmt"

	"github.com/michcald/files-organizer/internal"
)

func main() {
	fmt.Println("Hello")

	res := internal.Count(1, 2, 3)

	fmt.Println(res)
}
