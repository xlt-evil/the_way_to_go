package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("pdf_to_image/ePolicy.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
}
