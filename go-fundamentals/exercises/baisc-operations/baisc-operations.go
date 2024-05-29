package main

import "fmt"

func main() {
    var num1, num2 int
    fmt.Print("Ingresa el primer número: ")
    fmt.Scan(&num1)
    fmt.Print("Ingresa el segundo número: ")
    fmt.Scan(&num2)

    fmt.Printf("Suma: %d + %d = %d\n", num1, num2, num1+num2)
    fmt.Printf("Resta: %d - %d = %d\n", num1, num2, num1-num2)
    fmt.Printf("Multiplicación: %d * %d = %d\n", num1, num2, num1*num2)
    if num2 != 0 {
        fmt.Printf("División: %d / %d = %d\n", num1, num2, num1/num2)
    } else {
        fmt.Println("División: No se puede dividir por cero")
    }
}