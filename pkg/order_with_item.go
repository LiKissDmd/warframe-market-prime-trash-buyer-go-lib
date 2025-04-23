package warframe_market_prime_trash_buyer

import warframe_market_models "github.com/LiKissDmd/warframe-market-prime-trash-buyer-go-lib/internal/warframe_market/models"

type OrderWithItem struct {
	Order warframe_market_models.Order
	Item  warframe_market_models.ItemsItem
}
