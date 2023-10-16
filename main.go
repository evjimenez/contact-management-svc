package main

import "fmt"

// type Contact struct {
// 	name string
// 	tel  int
// 	mail string
// }

func main() {
	var num int
	menu()
	fmt.Scan(&num)
	switch num {
	case 1:
		addContact()
		break
	default:
		fmt.Println("You must to select a number between 1 and 6")
	}
}

func menu() {
	fmt.Println("---------Select an option---------")
	fmt.Println("1.Add Contact")
	fmt.Println("2.Display the list of existing contacts")
	fmt.Println("3.Search for a Contact by name")
	fmt.Println("4.Edit the information of existing contact")
	fmt.Println("5.Delet a contact from the list")
	fmt.Println("6.Exist the program")
	fmt.Print("\nYour choice: ")

}

//Agregar un nuevo contacto con la información de nombre, número de teléfono y correo electrónico.

func addContact() {
	var name string
	var mail string
	var tel int

	fmt.Println("Insert a name:")
	fmt.Scan(&name)

	fmt.Println("\nInsert an email:")
	fmt.Scan(&mail)

	fmt.Println("\nInsert a phone:")
	fmt.Scan(&tel)
}
