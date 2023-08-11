# 2023 noneland backend interview

## 前情提要

你現在在處理交易所風控後台的新 api 開設，以下有 `本次測驗需要新增的兩隻 api 規格`

並且有提供第三方 `XX交易所` api 文件的情況下

以下是 PM 提出的規格

## 基本規格

1. 第一隻 api，同時回傳兩種餘額
   1. 需要顯示 `XX交易所` 的 `現貨` 帳戶 USDT 餘額
   1. 需要顯示 `XX交易所` 的 `合約` 帳戶 USDT 餘額
2. 需要顯示 `現貨` 帳戶轉出轉入紀錄
   1. 根據法律遵循，我們應該保存 `6年` 內的所有交易紀錄

## 備註

可使用任何第三方工具、服務、套件來設計

## 加分題

1. 目前前端 app 的幣種報價使用的是 `現貨` 相關的 api，並且 api 存在呼叫限制，後台的呼叫不應該影響報價邏輯
2. 請撰寫可被測試的程式碼，或是直接附上測試程式
3. 架構也能調整，假設你覺得有更好的改法

## XX交易所 api 文件

### 現貨帳戶 api

#### api 使用限制

- REQUEST_WEIGHT 單位時間請求權重之和的上限
- RAW_REQUESTS 單位時間請求次數上限

GET `{{ url1 }}/exchangeInfo`

response:

```json
{
  "timezone": "UTC",
  "serverTime": 1565246363776,
  "rateLimits": [
     {
        "rateLimitType": "REQUEST_WEIGHT",
        "interval": "MINUTE",
        "intervalNum": 1,
        "limit": 1200
     },
     {
        "rateLimitType": "RAW_REQUESTS",
        "interval": "MINUTE",
        "intervalNum": 5,
        "limit": 6100
     }
  ]
}
```

#### 取得現貨帳戶餘額

GET `{{ url1 }}/spot/balance`

response:

```json
{
  "free": "10.12345"
}
```

#### 現貨帳戶轉入轉出紀錄

GET `{{ url1 }}/spot/transfer/records`

request 參數：

名稱 | 類型 | 是否必填 | 描述
----|------|--------|---------
startTime | LONG | NO |
endTime | LONG | NO |
endTime | LONG | NO |
current | LONG | NO | 當前回傳頁數，預設為 1
size | LONG | NO | 回傳筆數，預設 10，最大 100

response:

- status: PENDING (等待), CONFIRMED (成功), FAILED (失敗);

```json
{
   "rows": [
      {
         "amount": "0.10000000",
         "asset": "BNB",
         "status": "CONFIRMED",
         "timestamp": 1566898617,
         "txId": 5240372201,
         "type": "IN"
      },
      {
         "amount": "5.00000000",
         "asset": "USDT",
         "status": "CONFIRMED",
         "timestamp": 1566888436,
         "txId": 5239810406,
         "type": "OUT"
      },
      {
         "amount": "1.00000000",
         "asset": "EOS",
         "status": "CONFIRMED",
         "timestamp": 1566888403,
         "txId": 5239808703,
         "type": "IN"
      }
   ],
   "total": 3
}
```


### 合約帳戶 api

#### api 使用限制

- REQUEST_WEIGHT 單位時間請求權重之和的上限
- RAW_REQUESTS 單位時間請求次數上限

GET `{{ url2 }}/exchangeInfo`

response:

```json
{
  "timezone": "UTC",
  "serverTime": 1565246363776,
  "rateLimits": [
     {
        "rateLimitType": "REQUEST_WEIGHT",
        "interval": "MINUTE",
        "intervalNum": 1,
        "limit": 1200
     },
     {
        "rateLimitType": "RAW_REQUESTS",
        "interval": "MINUTE",
        "intervalNum": 5,
        "limit": 6100
     }
  ]
}
```


#### 取得合約帳戶餘額

GET `{{ url2 }}/futures/balance`

response:

```json
{
  "free": "10.12345"
}
```



