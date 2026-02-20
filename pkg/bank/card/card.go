package card

import "github.com/KadamovSh/DZ6A/pkg/bank/types"

// IssueCard создаёт карту
func IssueCard(currency types.Currency, color string, name string) types.Card {
    
    return types.Card{
		 ID:       1000,
        PAN:      "5058 xxxx xxxx 0001",
        Balance:  0,
        Currency: currency,  // берём из параметра
        Color:    color,      // берём из параметра
        Name:     name,       // берём из параметра
        Active:   true,
	}
}

//Withdraw - снимает деньги с карты
const withdrawLimit = 20_000_00

func Withdraw (card types.Card, amount types.Money) types.Card {
	if !card.Active{
		return card
	}

	if amount<0 {
		return card
	}

	if amount>withdrawLimit{
		return card
	}
	if amount> card.Balance{
		return card
	}

	card.Balance-=amount
	return card
}

