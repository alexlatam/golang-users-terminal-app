package main

import "fmt"
import "os"
import "bufio"
import "strings"

func crearUsuario() {
	fmt.Println("Crear usuario")
}

func listarUsuario() {
	fmt.Println("Listar usuarios")
}

func actualizarUsuario() {
	fmt.Println("Actualizar usuario")
}

func eliminarUsuario() {
	fmt.Println("Eliminar usuario")
}

func main() {


	// Esta variable es para leer lo que el usuario ingresa por consola
	var reader *bufio.Reader
	var option string
	
	reader = bufio.NewReader(os.Stdin)

	fmt.Println("A) Crear")
	fmt.Println("B) Listar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")

	fmt.Print("Ingrese una opción: ")
	// Leer lo que el usuario ingresa por consola, y se ira leyendo hasta que el usuario presione la tecla 'enter'	
	option, err := reader.ReadString('\n')

	if err != nil {
		panic("No es posible obtener el valor")
	}

	// Eliminar el salto de línea que se agrega al final de la cadena
	option = strings.TrimSuffix(option, "\n")

	switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuario()
		case "c", "actualizar":
			actualizarUsuario()
		case "d", "eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Opción no válida")
	}
}