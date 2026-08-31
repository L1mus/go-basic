package gobasic

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	helloWorld()
}

func helloWorld() {
	fmt.Println("hello world")
}
