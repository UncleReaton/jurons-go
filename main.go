package main

import (
	"fmt"
	"math/rand"
	"time"

	insults "jurons-go/assets"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	// ^ this one is depecrated, the new way is the one below
	rand.New(rand.NewSource(time.Now().UnixNano()))
	insultList := insults.GetInsult()
	count := len(insultList)
	random_line_number := rand.Intn(count) + 1
	fmt.Println(insultList[random_line_number])
}
