
package main

import "fmt"

func main() {
    costs := 250
    discountRate := 0.15
    discount := float64(costs) * discountRate

    fmt.Println("Original cost:", costs)
    fmt.Println("Discount:", discount)
    fmt.Println("Final price:", float64(costs)-discount)
}
