package helpers

import "fmt"

func CheckError(err error, msgErr string) {
	if err != nil {
		if msgErr != "" {
			fmt.Printf("[ERROR] : %s\n", msgErr)
			return
		}
		panic(err)
	}
}