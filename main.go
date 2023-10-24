package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name string
	tel  int
	mail string
}

func main() {
	var num int
	var choice string
	person := Person{}
	menu()
	fmt.Scan(&num)
	switch num {
	case 1:
		person.addContact()
		break
	default:
		fmt.Println("You must to select a number between 1 and 6")
		break
	}

	fmt.Println("Do you want to select another option: ")
	strings.Trim(choice, " ")
	fmt.Scan(&choice)

	if choice == "YES" {
		fmt.Println("What choice do you want: ")
		fmt.Scan(&num)
		switch num {
		case 1:
			person.addContact()
			break
		default:
			fmt.Println("You must to select a number between 1 and 6")
			break
		}
	} else {
		os.Exit(1)
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

func (p Person) addContact() {
	fmt.Println("Insert a name:")
	fmt.Scan(&p.name)

	fmt.Println("\nInsert an email:")
	fmt.Scan(&p.mail)

	fmt.Println("\nInsert a phone:")
	fmt.Scan(&p.tel)
}
