// package main

// import "fmt"

// func main() {
//     // Declaración de variables
//     var a int = 10  // Variable 'a' de tipo int con valor 10
//     var b = 20      // Variable 'b' con valor 20, tipo inferido como int
//     c := 30         // Variable 'c' con valor 30, tipo inferido como int

//     // Declaración de constantes
//     const pi = 3.14 // Constante 'pi' con valor 3.14

//     // Uso de variables y constantes
//     fmt.Println("a:", a)  // Imprime: a: 10
//     fmt.Println("b:", b)  // Imprime: b: 20
//     fmt.Println("c:", c)  // Imprime: c: 30
//     fmt.Println("pi:", pi) // Imprime: pi: 3.14
// }

// package main

// import "fmt"

// func main() {
//     // Uso de int para simplificar y dejar que el compilador maneje el tamaño
//     var edad int = 25
//     var salario int = 50000

//     fmt.Println("Edad:", edad)       // Imprime: Edad: 25
//     fmt.Println("Salario:", salario) // Imprime: Salario: 50000

//     // Uso de tamaños específicos cuando es necesario
//     var poblacionMundial int64 = 7800000000
//     fmt.Println("Población Mundial:", poblacionMundial) // Imprime: Población Mundial: 7800000000
// }

package main

import "fmt"

func main() {
    // Uso de int para la mayoría de los casos
    var edad int = 25
    var salario int = 50000

    fmt.Println("Edad:", edad)       // Imprime: Edad: 25
    fmt.Println("Salario:", salario) // Imprime: Salario: 50000

    // Uso de int64 para números muy grandes
    var poblacionMundial int64 = 7800000000
    fmt.Println("Población Mundial:", poblacionMundial) // Imprime: Población Mundial: 7800000000

    // Uso de int8 para optimización de memoria
    var temperatura int8 = 30
    fmt.Println("Temperatura:", temperatura) // Imprime: Temperatura: 30
}
