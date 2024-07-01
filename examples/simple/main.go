package main

import (
	"fmt"

	"github.com/pspiagicw/poll"
)

func main() {
	choice := poll.Choose("What do you want to eat?", []string{"Eat", "Sleep", "Code"})
	fmt.Println("You chose:", choice)

}
