package store

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WatchItem struct {
	Name          string    `bson:"name"`
	GoodsID       int       `bson:"goods_id"`
	TargetPrice   float64   `bson:"target_price"`
	AlertExpireAt time.Time `bson:"alert_expire_at"`
}

type PriceRecord struct {
	GoodsID   int       `bson:"goods_id"`
	Price     float64   `bson:"price"`
	CheckedAt time.Time `bson:"checked_at"`
}

var DB *mongo.Database

func InitMongo(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("buff")
}

func GetActiveWatchItems() []WatchItem {
	ctx := context.Background()
	now := time.Now()
	coll := DB.Collection("watch_items")
	cursor, err := coll.Find(ctx, bson.M{"alert_expire_at": bson.M{"$gt": now}})
	if err != nil {
		log.Println("find error:", err)
		return nil
	}
	var items []WatchItem
	cursor.All(ctx, &items)
	return items
}

func SavePriceRecord(goodsID int, price float64) {
	ctx := context.Background()
	record := PriceRecord{GoodsID: goodsID, Price: price, CheckedAt: time.Now()}
	_, err := DB.Collection("price_logs").InsertOne(ctx, record)
	if err != nil {
		log.Println("insert price error:", err)
	}
}

func SavePriceRecord2(goodsID int, price float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	doc := PriceRecord{
		GoodsID:   goodsID,
		Price:     price,
		CheckedAt: time.Now(),
	}
	result, err := DB.Collection("price_logs").InsertOne(ctx, doc)
	if err != nil {
		log.Println("插入价格失败:", err)
		return err
	}
	log.Printf("插入成功，ID=%v\n", result.InsertedID)
	return nil
}
