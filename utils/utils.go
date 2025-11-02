package utils

import (
	"fmt"
	"log"
	"strings"
)

func CheckError(e error, suc string, fai string) {
	if e != nil {
		if strings.Contains(e.Error(), "already exists") {
			fmt.Printf("%s\n", e)
		} else {
			log.Fatalf("Error %s:  %s", fai, e)
		}
	}
	fmt.Printf("Success: %s\n", suc)
}
