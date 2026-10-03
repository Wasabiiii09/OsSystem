package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welchen Ordner mochstest du erstellen ?")
	

	for {
		fmt.Print("Ordner Name: ")
		
		if !scanner.Scan(){
			break
		}

		OrderName := strings.TrimSpace(scanner.Text())

		if OrderName == "" {
			fmt.Println("Sie mussen ein Name Schreiben.")
			continue
		}

		_, err := os.Stat(OrderName)
		if err == nil {
			fmt.Println("Der Folder existiert schon")
			continue
		}

		errMkDirFolder := os.MkdirAll(OrderName, 0755)

		if errMkDirFolder != nil {
			fmt.Println("Fehler beim Erstellen bon Folder " + OrderName)
		} else {
			fmt.Printf("Ordner '%s', wurde erfolgreich erstellt.\n", OrderName)
		}

		

	}
}