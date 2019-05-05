package main

import (
	"crypto/sha256"
	"fmt"
)

// 散列算法不能解密
func main() {
	c := getSha256Code("18912345678")
	fmt.Println(c)
}

func getSha256Code(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
