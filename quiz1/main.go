package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	bossbaby "github.com/weeranieb/band-protocal-test/quiz1/bossBaby"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Receive first input
	fmt.Print("Enter input sequence of event today: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	// get the action "S and R" is input
	var action string
	fmt.Sscanf(input1, "%s", &action)

	// Call the GetBabyBehaviour function
	result := bossbaby.GetBabyBehaviour(action)
	fmt.Println(result)
}
