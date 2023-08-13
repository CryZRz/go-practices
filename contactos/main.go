package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"emial"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("Contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	errEncoder := encoder.Encode(contacts)

	if errEncoder != nil {
		return errEncoder
	}

	return nil

}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("Contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	errorDecoder := decoder.Decode(contacts)
	if errorDecoder != nil {
		return errorDecoder
	}

	return nil
}

func main() {
	var contacts []Contact

	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("error al cargar los contactos")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Gestor de contactos ====\n",
			"1. Agregar contacto \n",
			"2. Mostrar todos los contactos \n",
			"3. Salir\n",
			"Elige una opcion :)",
		)

		var option int
		_, err = fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Error al leer la opcion")
			return
		}

		switch option {
		case 1:
			var c Contact
			fmt.Println("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Println("Email")
			c.Email, _ = reader.ReadString('\n')
			fmt.Println("Telefono")
			c.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, c)

			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("error al guardar el contacto:", err)
			}
		case 2:
			fmt.Println("======================")
			for index, contact := range contacts {
				fmt.Printf("%d, Nombre: %s Email: %s Telefono %s \n", index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("======================")
		case 3:
			return
		default:
			fmt.Println("Opcion invalida")
			return
		}
	}
}
