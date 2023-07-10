package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// RotationalCipher shifts parameter 'plain' by 'shiftKey' characters
func RotationalCipher(plain string, shiftKey int) string {

	var abc = "abcdefghijklmnopqrstuvwxyz"
	var ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var cipher string
	var pos int

	if shiftKey%26 == 0 {
		// no rotation necessary
		return plain
	}

	for _, r := range strings.Split(plain, "") {

		if strings.Contains(abc, string(r)) {
			// found in lower case alphabet -> rotate

			pos = (strings.Index(abc, string(r)) + shiftKey) % 26
			cipher += string(abc[pos])

		} else if strings.Contains(ABC, string(r)) {
			// found in upper case alphabet -> rotate

			pos = (strings.Index(ABC, string(r)) + shiftKey) % 26
			cipher += string(ABC[pos])

		} else {
			// spaces and special characters -> leave as is

			cipher += string(r)
		}
	}

	return cipher

}

func main() {
	var rotation int
	var err error

	if (len(os.Args) > 1) && (len(os.Args[1]) > 0) {
		rotation, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Error during conversion of %s!", os.Args[1])
			return

		}

	} else {
		rotation = 0
	}

	stdin, err := ioutil.ReadAll(os.Stdin)

        if err != nil {
		log.Fatalf("Error reading from STDIN: %s!", err)		
	}
	str := string(stdin)


	fmt.Print(RotationalCipher(str, rotation))
}
