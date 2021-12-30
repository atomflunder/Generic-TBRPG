package utils

import (
	"fmt"
	"os"
)

func ExitGame() {
	fmt.Println("Thanks for playing. Hope to see you soon.")
	os.Exit(0)
}
