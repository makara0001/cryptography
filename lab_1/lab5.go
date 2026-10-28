package main
import(
	"fmt"
    "encoding/base64")

func Encode() string {
	var input string
	fmt.Printf("Enter text: ")
	fmt.Scanln(&input)

    // binary
	fmt.Printf("The binary: ")
	for i := 0; i < len(input); i++ {
		fmt.Printf("%08b ", input[i])
	}
	fmt.Println()

	// HEX 

	fmt.Printf("The Hex: ")
	for i := 0; i < len(input); i++ {
		fmt.Printf("%x", input[i])
	}
	fmt.Println()

	//base64

	encodedBase64 := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Printf("The Base64: %s\n", encodedBase64)



	return input
}

func main() {
	Encode()
}