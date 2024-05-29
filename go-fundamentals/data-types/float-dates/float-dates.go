package main

import "fmt"

func main() {
    // Inferencia automática a float64
    var a = 3.14
    fmt.Printf("a: %f, tipo: %T\n", a, a) // Imprime: a: 3.140000, tipo: float64

    // Especificación explícita a float32
    var b float32 = 3.14
    fmt.Printf("b: %f, tipo: %T\n", b, b) // Imprime: b: 3.140000, tipo: float32

    // Declaración corta infiere float64
    c := 3.14
    fmt.Printf("c: %f, tipo: %T\n", c, c) // Imprime: c: 3.140000, tipo: float64
}
