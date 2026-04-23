package hello

import (
	"fmt"

	"mateusf.com/greetings"
)

func main() {
	message := greetings.Hello("Mateus")
	fmt.Println(message)
}
