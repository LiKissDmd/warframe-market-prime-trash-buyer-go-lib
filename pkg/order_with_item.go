package warframe_market_prime_trash_buyer

import (
	"github.com/LiKissDmd/warframe-market-prime-trash-buyer-go-lib/internal/warframe_market/models"
)

// OrderWithItem связывает ордер с соответствующим айтемом для дальнейшей обработки.
type OrderWithItem struct {
	Order models.Order
	Item  models.ItemsItem
}
