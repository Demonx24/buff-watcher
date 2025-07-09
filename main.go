package main

import (
	"buff-watcher/resolver"
	"buff-watcher/scheduler"
	"buff-watcher/store"
)

func main() {
	store.InitMongo("mongodb://localhost:27017")
	resolver.GetPriceByGoodsID(39645)
	scheduler.StartPriceWatcher()
}
