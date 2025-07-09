package scheduler

import (
	"fmt"
	"time"

	"buff-watcher/notifier"
	"buff-watcher/resolver"
	"buff-watcher/store"
)

func StartPriceWatcher() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		items := store.GetActiveWatchItems()
		for _, item := range items {
			price, err := resolver.GetPriceByGoodsID(item.GoodsID)
			if err != nil {
				fmt.Println("价格获取失败:", err)
				continue
			}
			store.SavePriceRecord(item.GoodsID, price)
			if price <= item.TargetPrice {
				notifier.SendWeComAlert(item.Name, price)
			}
		}
	}
}
