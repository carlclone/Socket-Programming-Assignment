package common

import (
	"fmt"
	"os"
)

func PrintError(err error) {
	fmt.Println(err)
}

func CheckError(err error){
	if  err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}