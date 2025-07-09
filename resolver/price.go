package resolver

import (
	"buff-watcher/store"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPriceByGoodsID(goodsID int) (float64, error) {
	url := fmt.Sprintf("https://buff.163.com/api/market/goods/sell_order?game=csgo&goods_id=%d", goodsID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")
	fmt.Println(url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	items := result["data"].(map[string]interface{})["items"].([]interface{})
	if len(items) == 0 {
		return 0, fmt.Errorf("未找到价格")
	}
	priceStr := items[0].(map[string]interface{})["price"].(string)
	var price float64
	fmt.Sscanf(priceStr, "%f", &price)
	err = store.SavePriceRecord2(goodsID, price)
	if err != nil {
		fmt.Println("MongoDB保存价格失败:", err)
	}
	return price, nil
}
