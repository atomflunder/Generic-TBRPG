package utils

import (
	"fmt"
	"log"
	"os"
)

//deletes a file
func DeleteFile(s string) {
	err := os.Remove("./savedata/characters/" + s + ".json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully deleted profile of " + s)
}
