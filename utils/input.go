package utils

import (
	"bufio"
	"os"
	"strings"
)

//gets the user input
func GetUserInput() string {
	inpReader := bufio.NewReader(os.Stdin)
	inp, _ := inpReader.ReadString('\n')

	//we have to trim the whitespace from the user input
	inp = strings.TrimSpace(inp)

	return inp
}
