package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"heap/internal/heap"
	"heap/pkg/prettyprint"
)

func main() {
	mainMenuMsg := `Messing around with (max)heaps visualized™, 
  Please select your favorite operation:
    1 - Push Back
    2 - Push Front
    0 - Quit
`
	maxHeap := []int{50, 48, 45, 29, 15, 35, 40, 27, 26, 4, 12, 33, 30, 37, 20, 21, 19, 25}
	prompt := "> "
	askedToQuit := false
	reader := bufio.NewReader(os.Stdin)
	for !askedToQuit {
		clearTerminal()
		fmt.Print(mainMenuMsg)
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		option, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("ERROR: Failure while processing input, try again")
			continue
		}
		switch option {
		case 1:
			clearTerminal()
			fmt.Println(">>> Your (max)heap:")
			prettyprint.PrintHeap(maxHeap)
		input1:
			fmt.Println("What integer do you want to push back?")
			fmt.Print(prompt)
			input, _ = reader.ReadString('\n')
			value, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Println("ERROR: Failure while processing input, try again")
				goto input1
			}
			result := heap.InteractivelyPushBackMax(maxHeap, value)
			prettyprint.PrintHeap(result)
			fmt.Println(">>> Press <enter> to continue")
			fmt.Scanln()
		case 2:
			clearTerminal()
			fmt.Println(">>> Your (max)heap:")
			prettyprint.PrintHeap(maxHeap)
		input2:
			fmt.Println("What integer do you want to push front?")
			fmt.Print(prompt)
			input, _ = reader.ReadString('\n')
			value, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Println("ERROR: Failure while processing input, try again")
				goto input2
			}
			result := heap.InteractivelyPushFrontMax(maxHeap, value)
			prettyprint.PrintHeap(result)
			fmt.Print(">>> Press <enter> to continue")
			fmt.Scanln()
		case 0:
			fmt.Println("Bye!!!")
			os.Exit(0)
		default:
			fmt.Println("ERROR: Invalid Option, please try again")
		}
	}
}

func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func search(slice []int, value int) *int {
	for i, elem := range slice {
		if value == elem {
			return &i
		}
	}
	return nil
}
