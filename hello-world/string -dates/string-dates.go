package main

import "fmt"

func main() {
    // Declaración y asignación de una cadena (string)
    // Utilizado para representar secuencias de caracteres
    var saludo string = "Hola, Go!"

    // Las cadenas en Go son inmutables
    // Una vez creada, no se puede modificar
    // Para modificar una cadena, debemos crear una nueva
    // Por ejemplo, si intentamos cambiar el primer carácter, obtendremos un error:
    // saludo[0] = 'h' // Esto no es permitido en Go

    // UTF-8: Codificación utilizada para representar caracteres Unicode
    // Esto permite que las cadenas en Go contengan caracteres de muchos idiomas
	var mensaje string = "こんにちは世界" // "Hola Mundo" en japonés
	fmt.Println(mensaje)
    // Operaciones comunes con cadenas

    // Concatenación: Unir dos o más cadenas
    // Usamos el operador '+' para concatenar cadenas
    var nombre string = "Alice"
    var mensajeCompleto string = saludo + " " + nombre
    fmt.Println("Concatenación:", mensajeCompleto) // Imprime: Hola, Go! Alice

    // Longitud: Obtener la longitud de una cadena en bytes
    // La función len() devuelve la longitud de la cadena
    var longitud int = len(saludo)
    fmt.Println("Longitud:", longitud) // Imprime: 9 (número de bytes en "Hola, Go!")

    // Acceso a caracteres: Obtener un carácter específico de la cadena
    // Podemos acceder a un carácter utilizando su índice
    var primerCaracter byte = saludo[0]
    fmt.Printf("Primer carácter: %c\n", primerCaracter) // Imprime: Primer carácter: H

    // Subcadenas: Extraer una parte de la cadena
    // Usamos rebanadas (slices) para obtener una subcadena
    var subcadena string = saludo[0:4] // Obtiene los caracteres del índice 0 al 3
    fmt.Println("Subcadena:", subcadena) // Imprime: Hola

    // Otra operación común: Convertir una cadena a un array de bytes
    var arrayDeBytes []byte = []byte(saludo)
    fmt.Println("Array de bytes:", arrayDeBytes) // Imprime: Array de bytes: [72 111 108 97 44 32 71 111 33]

    // Convertir un array de bytes a una cadena
    var nuevaCadena string = string(arrayDeBytes)
    fmt.Println("Nueva cadena:", nuevaCadena) // Imprime: Nueva cadena: Hola, Go!
}
