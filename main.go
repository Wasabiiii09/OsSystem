package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"log"
)

func main() {
	
	scanner := bufio.NewScanner(os.Stdin)

	for { 	

		now := time.Now()

		formatted := now.Format("02.01.2006 15:04:05")

		file, errlog := os.OpenFile(
			"app.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)

		if errlog != nil {
			log.Fatal(errlog)
		}

		defer file.Close()

		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()

		fmt.Println("was Mocheten sie machen: ")
		fmt.Println("0.Exit")
		fmt.Println("1.Folder erstellen")
		fmt.Println("2.Folder Loschen")

		fmt.Printf("#: ")
		if !scanner.Scan() {
    		break
		}
		UserInput := strings.TrimSpace(scanner.Text())

		

		if UserInput == "1" {

			clear = exec.Command("clear")
			clear.Stdout = os.Stdout
			clear.Run()

    	    fmt.Println(" Welchen Ordner mochstest du erstellen ?")
			fmt.Println("(Schreiben sie 0 um dem Function zu verlassen)")
    	    
			for {
    	    	
				fmt.Printf("#: ")
    	    	if !scanner.Scan(){
    	    		break
    	    	}
				
    	    	OrderName := strings.TrimSpace(scanner.Text())
        
    	    	if OrderName == "" {
    	    		fmt.Println("Sie mussen ein Name Schreiben.")
    	    		continue
    	    	}
				
				if OrderName == "0" {
					break
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
					file.WriteString( formatted + " - Folder " + OrderName + " erstellt \n")
    	    	}

		    }	
		}

		if UserInput == "2" {
			
			clear = exec.Command("clear")
			clear.Stdout = os.Stdout
			clear.Run()

			fmt.Println(" Schreiben sie die name des Folders was sie Loschen wollen")
			fmt.Println("(Schreiben sie 0 um den Function zu beendden)")

			for {

				fmt.Printf("#: ")
				if !scanner.Scan() {
					break
				}

				FolderNameDelete := strings.TrimSpace(scanner.Text())
				
				if FolderNameDelete == "" {
					fmt.Println("Sie mussen ein Name von Folder Schreiben.")
					continue
				}

				if FolderNameDelete == "0" {
					break
				}

				_, err := os.Stat(FolderNameDelete)

				if err != nil {
					fmt.Println("Solchen Folder existiert nicht.")
					continue
				}

				fmt.Printf("Sind sie sicher das sie den Folder '[%s]' loschen wollen? (y/n): ", FolderNameDelete)
				
				if !scanner.Scan() {
    				break
				}

				SicherLoschen := strings.TrimSpace(scanner.Text())

				if SicherLoschen == "y" {
					
					err = os.RemoveAll(FolderNameDelete)
					if err != nil {
						fmt.Println("Error beim Folder Loschen")
						continue
					}

					fmt.Printf("Der File '%s' wurde erfolgreich geloscht", FolderNameDelete)
					time.Sleep(3 * time.Second)

					file.WriteString( formatted + " - Folder " + FolderNameDelete + " geloscht \n")
					
					clear = exec.Command("clear")
			        clear.Stdout = os.Stdout
			        clear.Run()
        
			        fmt.Println(" Schreiben sie die name des Folders was sie Loschen wollen")
			        fmt.Println("(Schreiben sie 0 um den Function zu beendden)")

					continue

				} else if SicherLoschen == "n" {
					
					fmt.Printf("Der Folder '%s' wird nicht geloscht", FolderNameDelete)
					continue

				} else { 
					clear = exec.Command("clear")
			        clear.Stdout = os.Stdout
			        clear.Run()

					fmt.Println("Schreibem sie y oder n")
					fmt.Println("Folder Loschen wurde abgebrochen")

					time.Sleep(2 * time.Second)

					clear = exec.Command("clear")
			        clear.Stdout = os.Stdout
			        clear.Run()
        
			        fmt.Println(" Schreiben sie die name des Folders was sie Loschen wollen")
			        fmt.Println("(Schreiben sie 0 um den Function zu beendden)")

					continue
				}
			}
		}
    }	
}