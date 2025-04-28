package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
)

func main() {
	//GMT time zone
	gmt, _ := time.LoadLocation("Europe/Dublin")

	currentTime := time.Now().In(gmt).Format("2006-01-02 15:04:05")
	fmt.Println(currentTime)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	qrData := name + " " + currentTime

	qrErr := qrcode.WriteFile(qrData, qrcode.Medium, 256, "qrcode1.png")
	if qrErr != nil {
		fmt.Println("Error creating QR code: ", qrErr)
		return
	}

	fmt.Println("Successfully generated")
}
