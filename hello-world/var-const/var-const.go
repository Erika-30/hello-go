package main

import "fmt"

func main() {
    // Declaración de variables
    var a int = 5  // Variable 'a' de tipo int con valor 5
    var b int = 30 // Variable 'b' de tipo int con valor 30
    c := 12        // Variable 'c' con valor 12, tipo INFERIDO como int

    // Declaración de constantes
    const pi = 3.14159 // Constante 'pi' con valor 3.14159

    // Uso de fmt.Println para imprimir valores
    fmt.Println("a:", a)  // Imprime: a: 5
    fmt.Println("b:", b)  // Imprime: b: 30
    fmt.Println("c:", c)  // Imprime: c: 12
    fmt.Println("pi:", pi) // Imprime: pi: 3.14159
}
