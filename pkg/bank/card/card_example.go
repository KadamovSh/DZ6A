package card

import (
	"fmt"

	
	"github.com/KadamovSh/DZ6A/pkg/bank/types"
)

func ExampleWithdraw_positive() {
    card := types.Card{
        Balance: 30_000_00,
        Active:  true,
    }
    result := Withdraw(card, 10_000_00)
    fmt.Println(result.Balance)
    // Output: 2000000
}

func ExampleWithdraw_noMoney () {
	card := types.Card{
		Balance: 10_000_00,
		Active: true ,
	}
	result := Withdraw(card, 20_000_00) 
	fmt.Println(result)
	// Output: 1000000
}

func ExampleWithdraw_inactive() {
    card := types.Card{
        Balance: 30_000_00,
        Active:  false,
    }
    result := Withdraw(card, 10_000_00)
    fmt.Println(result.Balance)
    // Output: 3000000
}

func ExampleWithdraw_limit() {
    card := types.Card{
        Balance: 50_000_00,
        Active:  true,
    }
    result := Withdraw(card, 30_000_00)
    fmt.Println(result.Balance)
    // Output: 5000000
}
