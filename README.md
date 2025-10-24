# go-routine

## 執行方式

1. 進到專案資料夾
```bash
cd go-routine
```

2. 啟動 server
```bash
go run ./cmd/server
```

3. 呼叫 API
```bash
curl -X POST http://localhost:8080/login           -H "Content-Type: application/json"           -d '{"userId":"user123","password":"pw123"}'
```

4. 成功回應範例
```json
{
  "token": "token-user123-1730xxxxxxxxxxxxxx",
  "dashboard": {
    "username": "Bruce",
    "balance": 12345.67,
    "cards": [
      { "number": "****-****-****-1234", "brand": "VISA" },
      { "number": "****-****-****-5678", "brand": "Mastercard" }
    ]
  }
}
```

## 說明

- `dashboard_service.go` 用 `sync.WaitGroup` 同時去抓，可以減少延遲
  - 信用卡清單 (card service)...Sleep:1500ms
  - 帳戶餘額 (account service)...Sleep:1200ms
  - 使用者名稱 (DB via repo)...Sleep:1000ms
