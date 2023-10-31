package jwtX

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(2, "xyq")
	fmt.Println("token:", token)
	fmt.Println("err:", err)
}
func TestA(t *testing.T) {
}
