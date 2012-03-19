package main

import (
	"fmt"
	"bitbucket.org/jzs/quizletongo"
)

func main() {
	fmt.Printf("Here's a set from Quizlet\n")

	session := quizongo.NewSession("apikey")
	set, err := session.Set(4797490, nil)

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("Set Id: %v\n", set.Id)
	fmt.Printf("Title: %v\n", set.Title)
}
