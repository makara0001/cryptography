package main
import"fmt"

func xorEncrypt(text string, key byte) string {

	input := []byte(text)

	for i := 0; i < len(input); i++{
		input[i] = input[i] ^ key

	} 
	
	return string (input)

}
func main(){
	var text string
	var keys string

	fmt.Printf("Enter text: ")
    fmt.Scanln(&text)

	fmt.Printf("Enter key character: ")
    fmt.Scanln(&keys)

	var key byte = keys[0]

	encrypted := xorEncrypt(text, key)
    fmt.Printf("Result: ")
    for i := 0; i < len(encrypted); i++ {
        fmt.Printf("%x ", encrypted[i])
    }
    fmt.Println()

	decrypted := xorEncrypt(encrypted, key)
	fmt.Println("Decrypted:", decrypted)


}