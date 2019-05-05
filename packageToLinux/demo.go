package main

import (
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("/bin/bash", "-c", "asciiview", "瓜皮.jpg")
	f, _ := os.Create("my.txt")
	c.Stdout = f
	c.Run()
}
