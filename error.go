package main

import (
	"fmt"
	"log"
)

func CheckErr(err error, msg string, args ...any) {
	if err != nil {
		log.Fatal(fmt.Sprintf("- "+msg+"\n"+err.Error(), args...))
	}
}
