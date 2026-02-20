package main

import (
    "fmt"
    "github.com/KadamovSh/DZ6A/pkg/bank/card"
    "github.com/KadamovSh/DZ6A/pkg/bank/types"
)


func main() {
	myCard:=card.IssueCard(types.TJS, "Black", "Shoir Qadamov")
	fmt.Println(myCard)
}