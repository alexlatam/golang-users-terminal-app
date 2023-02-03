package main

import (
	"fmt"
 	"os"
 	"bufio"
 	"strings"
 	"strconv"
 	"os/exec"
)

// Esta variable es para leer lo que el usuario ingresa por consola
var reader *bufio.Reader

// Definir la estructura de un usuario. La 'clase' User
type User struct {
	id int
	name string
	email string
	age int
}

var id int = 0
// Definir un mapa de usuarios. Que contendrá todos los usuarios que se creen
var users = make(map[int]User)

func menu() {
	fmt.Println("A) Crear")
	fmt.Println("B) Listar")
	fmt.Println("C) Actualizar")
	fmt.Println("D) Eliminar")
}

func crearUsuario() {
	clearConsole()
	
	fmt.Print("Ingrese el nombre de usuario: ")
	name := readLine()

	fmt.Print("Ingrese el email de usuario: ")
	email := readLine()

	fmt.Print("Ingrese la edad de usuario: ")
	age, err := strconv.Atoi(readLine()) // Convertir el string a entero

	if err != nil {
		panic("No es posible convertir la edad a entero")
	}

	// Incrementamos el id
	id++
	// Creamos un nuevo usuario
	user := User { id, name, email, age }
	// Agregamos el usuario al mapa de usuarios
	// Recordar: un mapa es como un objeto literal en JS
	users[id] = user

	fmt.Println(">>> Usuario creado exitosamente!\n")
}

func listarUsuario() {

	clearConsole()
	fmt.Println(">>> Usuarios <<<")
	for id, user := range users {
		fmt.Println(id, "-", "Nombre:", user.name, ", Email:", user.email, ", Edad:", user.age, "años")
	}
	fmt.Println("\n")
}

func actualizarUsuario() {
	clearConsole()
	fmt.Print("Ingrese el id del usuario a editar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el id a entero")
	}

	if _, ok := users[id]; !ok {
		panic(">>> No existe un usuario con el id proporcionado")
	}

	fmt.Print("Ingrese el nombre de usuario: ")
	name := readLine()

	fmt.Print("Ingrese el email de usuario: ")
	email := readLine()

	fmt.Print("Ingrese la edad de usuario: ")
	age, err := strconv.Atoi(readLine()) // Convertir el string a entero

	if err != nil {
		panic("No es posible convertir la edad a entero")
	}

	user := User { id, name, email, age }
	users[id] = user

	fmt.Println(">>> Usuario actualizado exitosamente!\n")

}

func eliminarUsuario() {

	clearConsole()
	fmt.Print("Ingrese el id del usuario a eliminar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No es posible convertir el id a entero")
	}

	if _, ok := users[id]; !ok {
		panic(">>> No existe un usuario con el id proporcionado")
	}

	delete(users, id)
	fmt.Println(">>> Usuario eliminado exitosamente!\n")

}

// Funcion que ejecuta un comando de consola para limpiar la pantalla
// De esta manera se puede ejecutar cualquier comando de consola desde Go
func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {
	
	// Leer lo que el usuario ingresa por consola, y se ira leyendo hasta que el usuario presione la tecla 'enter'	
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		// Eliminar el salto de línea que se agrega al final de la cadena
		return strings.TrimSpace(option)
	}
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Ups! Ocurrió un error!")
		}
	}()
	
	var option string
	
	reader = bufio.NewReader(os.Stdin)

	for {

		menu()
	
		fmt.Print("Ingrese una opción: ")
		option = readLine()
	
		if option == "quit" || option == "q" {
			break
		}
	
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

	fmt.Println("Adios")

}