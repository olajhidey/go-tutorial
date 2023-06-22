package main

import (
	"fmt"

	"github.com/olajhidey/hello/utils"
)

func main() {

	work := make(chan string, 3)
	fin := make(chan string)

	go utils.ChannelCloseTutorial(work, fin)

	word := "cloudacademy"

	for j := 0; j < len(word); j++ {
		letter := string(word[j])
		work <- letter
		fmt.Printf("%s sent ...\n", letter)
	}

	close(work)

	fmt.Printf("result: %s", <-fin)

}
