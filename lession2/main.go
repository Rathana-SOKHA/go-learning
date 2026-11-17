package main

import "fmt"

func main() {
    length := 8
    width := 5

    area := length * width
    perimeter := 2 * (length + width)

    fmt.Println("Length:", length)
    fmt.Println("Width:", width)
    fmt.Println("Area:", area)
    fmt.Println("Perimeter:", perimeter)
}