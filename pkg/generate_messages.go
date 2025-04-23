package warframe_market_prime_trash_buyer

import (
	"fmt"
)

func GeneratePurchaseMessages(profitableOrders []OrderWithItem) ([]string, error) {
	messages := make([]string, 0, len(profitableOrders))
	for _, o := range profitableOrders {
		user := o.Order.User.IngameName
		item := o.Item.ItemName
		price := min(3, o.Order.Platinum)
		qty := o.Order.Quantity
		total := price * qty

		msg := fmt.Sprintf(
			"/w %s Hi, %s! You have WTS order: %s for %d. I would like to buy all %d for %d if you are interested :)",
			user, user, item, o.Order.Platinum, qty, total,
		)
		messages = append(messages, msg)
	}
	return messages, nil
}
