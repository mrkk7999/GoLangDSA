package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var strBuffer bytes.Buffer
	strBuffer.WriteString("Kiran")
	strBuffer.WriteString(" Kshirsagar")
	fmt.Println("The string buffer output is", strBuffer.String())

	var strBuffer1 bytes.Buffer
	fmt.Fprintf(&strBuffer1, "It is number : %d, This is a string : %v", 79, "seventy nine")
	strBuffer1.WriteString("[Done]")
	fmt.Println("The string buffer output is", strBuffer1.String())

	var byteString bytes.Buffer
	byteString.Write([]byte("Writting"))
	fmt.Fprintf(&byteString, " Inside bytestring")
	byteString.WriteTo(os.Stdout)

	var strByyt bytes.Buffer
	strByyt.Grow(64)
	strByytestrbyyte := strByyt.Bytes()
	strByyt.Write([]byte("It is a 64 byte"))
	fmt.Printf("%b", strByytestrbyyte[:strByyt.Len()])
}
