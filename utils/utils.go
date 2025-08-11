package utils

import (
	"fmt"
	"log"
	"strings"
)

func CheckError(e error, s string) {
	if e != nil {
		if strings.Contains(e.Error(), "already exists") {
			fmt.Printf("%s", e)
		} else {
			log.Fatalf("Error %s:  %s", s, e)
		}
	}
}
