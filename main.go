package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"time"
)

var (
	houses = [4]string{
		"Gryffindor",
		"Hufflepuff",
		"Ravenclaw",
		"Slytherin",
	}
	bubbles = [6]string{
		"Let me see...",
		"Hm, I think...",
		"Well, well...",
		"No question about it!",
		"This is easy.",
		"Of course!",
	}
)

func main() {

	user, err := user.Current()
	if err != nil {
		fmt.Println("I'm sorry, I cant decide.")
		os.Exit(1)
	}

	var sum int64
	for _, val := range []byte(user.Username) {
		sum += int64(val)
	}

	rand.Seed(time.Now().UnixNano())
	bubble := bubbles[rand.Intn(len(bubbles))]

	rand.Seed(sum % int64(len(houses)))
	house := houses[rand.Intn(len(houses))]

	output := fmt.Sprintf("%s You are a %s!", bubble, house)

	typeAsBot(output)
}

func typeAsBot(input string) {
	rand.Seed(time.Now().UnixNano())

	for _, l := range input {
		time.Sleep(botDelay(10, 5, 15))
		fmt.Printf("%s", string(l))
	}
	fmt.Println()
}

func botDelay(base int, delta int, multiplier int) time.Duration {
	return time.Millisecond * (time.Duration(base) + time.Duration(rand.Intn(delta)*multiplier))
}
