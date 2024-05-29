package main

import (
	"fmt"
	"math"
)

func main() {
    const pi = 3.14159
    var radio float64

    fmt.Print("Ingresa el radio del círculo: ")
    fmt.Scan(&radio)

    area := pi * math.Pow(radio, 2)
    fmt.Printf("El área del círculo con radio %.2f es %.2f\n", radio, area)
}