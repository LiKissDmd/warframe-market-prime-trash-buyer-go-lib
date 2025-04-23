package warframe_market_prime_trash_buyer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	warframe_market "github.com/LiKissDmd/warframe-market-prime-trash-buyer-go-lib/internal/warframe_market"
	warframe_market_models "github.com/LiKissDmd/warframe-market-prime-trash-buyer-go-lib/internal/warframe_market/models"
)

func GetProfitableOrders() ([]OrderWithItem, error) {
	items, err := fetchItems()
	if err != nil {
		return nil, err
	}

	itemsToBuy := filterItems(items, ProfitableItemNames)
	profitableOrders := make([]OrderWithItem, 0, 10)
	rateLimiter := time.NewTicker(time.Second / warframe_market.MaxRequestsPerSecond)
	defer rateLimiter.Stop()

	for _, item := range itemsToBuy {
		<-rateLimiter.C

		orders, err := fetchOrders(item.URLName)
		if err != nil {
			return nil, fmt.Errorf("error getting orders for %s: %w", item.ItemName, err)
		}

		for _, order := range orders {
			if isOrderProfitable(order) {
				profitableOrders = append(profitableOrders, OrderWithItem{
					Order: order,
					Item:  item,
				})
			}
		}
	}

	return profitableOrders, nil
}

func fetchItems() ([]warframe_market_models.ItemsItem, error) {
	resp, err := http.Get("https://api.warframe.market/v1/items")
	if err != nil {
		return nil, fmt.Errorf("failed to get items info: %w", err)
	}
	defer resp.Body.Close()

	var result warframe_market_models.ItemsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return result.Payload.Items, nil
}

func fetchOrders(itemURLName string) ([]warframe_market_models.Order, error) {
	url := fmt.Sprintf("https://api.warframe.market/v1/items/%s/orders", itemURLName)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result warframe_market_models.OrdersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Payload.Orders, nil
}

func filterItems(items []warframe_market_models.ItemsItem, names []string) []warframe_market_models.ItemsItem {
	var result []warframe_market_models.ItemsItem
	for _, item := range items {
		if slices.Contains(names, item.ItemName) {
			result = append(result, item)
		}
	}
	return result
}

func isOrderProfitable(order warframe_market_models.Order) bool {
	return order.OrderType.IsSell() &&
		order.User.Status.IsIngame() &&
		order.Platform.IsPC() &&
		order.Quantity >= 3 &&
		order.Platinum <= 4
}
