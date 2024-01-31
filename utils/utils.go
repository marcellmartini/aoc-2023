package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}
