package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	rescuechicken "github.com/weeranieb/band-protocal-test/quiz2/rescueChicken"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Receive first input
	fmt.Print("Enter input chicken No and length of roof: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	// Receive second input
	fmt.Print("Enter input chicken position: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	// get the chickenNo and length when "x y" is input
	var chickenNo, length int
	fmt.Sscanf(input1, "%d %d", &chickenNo, &length)

	// get the position of the chickens when "a b c" is input
	var positions []int
	parts := strings.Split(input2, " ")

	for _, part := range parts {
		// Convert each part to an integer
		if value, err := strconv.Atoi(part); err == nil {
			positions = append(positions, value)
		} else {
			fmt.Println("Error converting string to integer:", err)
		}
	}
	positions = positions[:chickenNo]

	// Call the RescueChicken function
	result := rescuechicken.RescueChicken(chickenNo, length, positions)
	fmt.Println(result)
}
