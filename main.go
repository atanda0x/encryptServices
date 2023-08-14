package main

import (
	"log"

	"github.com/atanda0x/encryptString/utils"
)

func main() {
	key := "111023043350789514532147"
	message := "I'm a programmer"
	// log.Println("Original message ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted message: ", encryptedString)
	decryptedString := utils.DecrptString(key, encryptedString)
	log.Println("Decrypted message: ", decryptedString)
}
