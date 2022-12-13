package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(ComData string, allnotes []string) {

	for {
		fmt.Println("Enter a command and data:")
		UserInput := bufio.NewScanner(os.Stdin)
		UserInput.Scan()
		ComData = UserInput.Text()
		Com___Data := strings.SplitN(ComData, " ", 2)

		switch Com___Data[0] {
		case "exit":
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		case "create":
			create(Com___Data, allnotes)
		case "list":
			list(allnotes)
		case "clear":
			clear(allnotes)
		case "update":
			update(Com___Data, allnotes)
		case "delete":
			delete(Com___Data, allnotes)
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func create(Com___Data []string, allnotes []string){
	for i, el := range allnotes{
		if len(Com___Data) < 2 {
			fmt.Println("[Error] Missing note argument")
			break
		} else if el == "" {
			allnotes[i] = Com___Data[1]
			fmt.Println("[OK] The note was successfully created")
			break
		} else if Com___Data[1] == " "{
			fmt.Println("[Error] Missing note argument")
			break
		} else if i+1 == len(allnotes) {
			fmt.Println("[Error] Notepad is full")
		}
	}
}

func list (allnotes []string) {
	counter:=0
	for j:= 0; j < len(allnotes); j++ {
		if allnotes[j] == "" {
			counter += 1
		}
	}
	if counter == len(allnotes) {
		fmt.Println("[Info] Notepad is empty")
	}

	for i, el := range allnotes {
		if el != "" {
			fmt.Printf("[Info] %d: %s\n", i+1, el)
		}
	}
}


func clear (allnotes []string){
	for i, _ := range allnotes {
		allnotes[i] = ""
	}
	fmt.Println("[OK] All notes were successfully deleted")
}

func update(Com___Data []string, allnotes []string){
	if len(Com___Data) < 2 {
		fmt.Println("[Error] Missing position argument")
	} else {
		Up := strings.SplitN(Com___Data[1], " ", 2)
		if len(Up) < 2 {
		fmt.Println("[Error] Missing note argument")
		} else {
			position, err := strconv.Atoi(Up[0])
			if err != nil {
				fmt.Println("[Error] Invalid position:", Up[0])
			} else if position < 1 || position > len(allnotes) {
				fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, len(allnotes))
			} else if allnotes[position-1] == "" {
				fmt.Println("[Error] There is nothing to update")
			} else {
				allnotes[position-1] = Up[1]
				fmt.Printf("[OK] The note at position %v was successfully updated\n", Up[0])
			}
		}
	}
}

func delete(Com___Data []string, allnotes []string){
	if len(Com___Data) < 2 {
		fmt.Println("[Error] Missing position argument")
	} else {
		Del := Com___Data[1]
		if Del == "" {
		fmt.Println("[Error] Missing position argument")
			} else {
			position, err := strconv.Atoi(Del)
			if err != nil {
				fmt.Println("[Error] Invalid position:", Del)
			} else if position < 1 || position > len(allnotes) {
				fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, len(allnotes))
			} else if allnotes[position-1] == "" {
				fmt.Println("[Error] There is nothing to delete")
			} else {
				allnotes = append(allnotes[:position-1], allnotes[position:]...)
				allnotes = append(allnotes, "")
				fmt.Printf("[OK] The note at position %v was successfully deleted\n", Del)
			}
		}
	}
}

func main() {
	var numbers int
	fmt.Println("Enter the maximum number of notes:")
	fmt.Scan(&numbers)

	var ComData string
	var allnotes [] string
	allnotes = make ([] string, numbers)

	input (ComData, allnotes)
}