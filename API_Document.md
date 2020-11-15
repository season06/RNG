# RNG web API Document

## Draw Number
Description: 抽出範圍為 min <= i <= max 的數字, 並設定數量
```
URL: /draw_number
Method: POST
Header:
{
	"Content-Type": "application/json"
}
Body:
{
    "Amount": 3,
    "Min": 1,
    "Max": 3
}
```
```
Backend_Response:
Status Code: 200
{
    "result": [3, 1, 1]
}

Status Code: 400
Situation: Amount is empty
{
    "result": "Enter the amount"
}

Status Code: 400
Situation: Max value is empty
{
    "result": "Enter a max number"
}

Status Code: 400
Situation: Min value > Max value
{
    "result": "Min value musts be smaller than Max value"
}
```

## Draw Items
Description: 輸入抽獎的獎品, 並抽出一個
```
URL: /draw_items
Method: POST
Header:
{
	"Content-Type": "application/json"
}
Body:
{
    "Item": ["A", "B", "C", "D"]
}
```
```
Backend_Response:
Status Code: 200
{
    "result": "A"
}

Status Code: 400
Situation: Items is empty
{
    "result": "Enter the Items"
}
```

## Draw
Description: 抽出中獎人與中獎獎品的對應, 並可設定獎品數量
```
Url: /draw_number
Method: POST
Header:
{
	"Content-Type": "application/json"
}
Body:
{
    "Award": ["A", "B"],
    "Award_amount": [2, 2],
    "Candidate": ["I", "He", "She"]
}
```
```
Backend_Response:
Status Code: 200
[
    {
        "award": "A",
        "candidate": "He"
    },
    {
        "award": "A",
        "candidate": "She"
    },
    {
        "award": "B",
        "candidate": "I"
    }
]

Status Code: 400
Situation: Items is empty
{
    "result": "Enter the Items"
}

Status Code: 400
Situation: People is empty
{
    "result": "Enter the People"
}

Status Code: 400
Situation: 獎品數量 > 抽獎人數量
{
    "result": "Amount of Award is larger then Amount of People"
}
```