```csharp
buff-watcher/
├── cmd/
│   ├── search.go        # 主动查询饰品信息
│   └── add.go           # 添加新饰品到配置并解析 goods_id
├── config/
│   └── watchlist.json   # 配置文件（自动维护）
├── resolver/
│   └── goodsid.go       # 名称转 goods_id 查询工具
├── scheduler/
│   └── cron.go          # 定时轮询器
├── store/
│   └── mongo.go         # MongoDB 存取
├── notifier/
│   └── wecom.go         # 企业微信通知模块
├── utils/
│   └── http.go          # 通用请求封装
├── main.go              # 启动服务入口

```

