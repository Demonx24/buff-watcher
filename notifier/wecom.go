package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func SendWeComAlert(name string, price float64) error {
	webhook := os.Getenv("WECOM_WEBHOOK")
	content := fmt.Sprintf("\u26a0\ufe0f [%s] 当前价格 %.2f 已低于预设\nhttps://buff.163.com/goods/%s", name, price, name)
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": content,
		},
	}
	data, _ := json.Marshal(payload)
	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
