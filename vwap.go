package main

import (
    "fmt"
    "log"
    "net/url"
    "encoding/json"
    "github.com/gorilla/websocket"
    "strings"
    "strconv"
)

var (
	multiline = `
	{
    "type": "subscribe",
    "product_ids": [
        "BTC-USD",
        "ETH-USD",
	"ETH-BTC"
    ],
    "channels": [
        "matches",
        {
            "product_ids": [
                "BTC-USD",
                "ETH-USD",
		"ETH-BTC"
            ]
        }
    ]
}
	`
	subscribe  = []byte(multiline)
)


func main() {

//    type Employee struct {
 //	 Name, City string
//	 Salary int32
//	}

   type Match struct {
	 Type string
	 TradeId int32
	 MakerOrderId string
	 TakerOrderId string
	 Side string
	 Size string
	 Price string
	 ProductId string `json:"product_id"`
	 Sequence int32
	 Time string
 }

 var match Match

/*  "type": "match",
  "trade_id": 137835913,
  "maker_order_id": "3d93e427-e0ea-4a9d-80c4-e9022566d356",
  "taker_order_id": "d9af72bf-a80d-4921-a077-0dfe807785c0",
  "side": "sell",
  "size": "0.0769",
  "price": "1873.78",
  "product_id": "ETH-USD",
  "sequence": 18979534058,
  "time": "2021-07-16T09:55:23.529193Z"
*/

    u, err := url.Parse("wss://ws-feed.pro.coinbase.com:443")
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        // handle error
    }
    // send message
    c.WriteMessage(websocket.TextMessage, subscribe)
    // receive message

    _, message, err := c.ReadMessage()
    if err != nil {
	  log.Println(err)
  }

    // Print initial message
    fmt.Println(string(message))

    var BtcUsdVol [200]float64 = [200]float64{} // Partial assignment
    var BtcUsdHigh float64 = 0
    var BtcUsdLow float64 = 0
    var BtcUsdCheck float64 = 0
    var BtcUsdTP float64 = 0
    var BtcUsdCum float64 = 0
    var BtcUsdCnt int = 0
    var BtcUsdSize float64 = 0

    var EthUsdVol [200]float64 = [200]float64{} // Partial assignment
    var EthUsdHigh float64 = 0
    var EthUsdLow float64 = 0
    var EthUsdCheck float64 = 0
    var EthUsdTP float64 = 0
    var EthUsdCum float64 = 0
    var EthUsdCnt int = 0
    var EthUsdSize float64 = 0

    var EthBtcVol [200]float64 = [200]float64{} // Partial assignment
    var EthBtcHigh float64 = 0
    var EthBtcLow float64 = 0
    var EthBtcCheck float64 = 0
    var EthBtcTP float64 = 0
    var EthBtcCum float64 = 0
    var EthBtcCnt int = 0
    var EthBtcSize float64 = 0



    for 1 == 1 {
    _, message, err := c.ReadMessage()
    if err != nil {
	  log.Println(err)
    }

//    myString := string(message)
//    fmt.Println(myString)
    json.Unmarshal([]byte(string(message)), &match)
//  Ignore all sell matches, only focus on buy
    if strings.Contains(match.Side,"buy") {
    fmt.Printf("%s ",match.ProductId)
    //
    //
    // Figure out for BTC-USD
    if strings.Contains(match.ProductId,"BTC-USD") {
        BtcUsdCheck,err = strconv.ParseFloat(match.Price, 64)
        // Find High value
        if BtcUsdCheck > BtcUsdHigh { BtcUsdHigh = BtcUsdCheck }
        // Find Low value
        if BtcUsdLow == 0 { BtcUsdLow = BtcUsdHigh }
        if BtcUsdCheck < BtcUsdLow { BtcUsdLow = BtcUsdCheck }
    // Fudge Typical Price
    // https://corporatefinanceinstitute.com/resources/knowledge/trading-investing/volume-weighted-adjusted-price-vwap/
    BtcUsdTP = (BtcUsdHigh + BtcUsdLow + BtcUsdCheck) / 3
    // Fudge Cumulative Volume by storing last 200 in array
    // Ignore the fact that many are zero size at first
    // Use Modulus to continue to reuse array after hitting 200 record mark
    if BtcUsdCnt < 200 {
    BtcUsdVol[BtcUsdCnt],err = strconv.ParseFloat(match.Size, 64)
    } else  {
    BtcUsdVol[BtcUsdCnt % 200],err = strconv.ParseFloat(match.Size, 64)
    }

    for _, value := range BtcUsdVol {
		BtcUsdCum = BtcUsdCum + value
	}

    BtcUsdSize,err = strconv.ParseFloat(match.Size,64)
    fmt.Printf(" %f",BtcUsdTP*BtcUsdSize/BtcUsdCum)

    fmt.Println()
    BtcUsdCnt += 1
    } // End of BTC-USD
        // Figure out for ETH-USD
    if strings.Contains(match.ProductId,"ETH-USD") {
        EthUsdCheck,err = strconv.ParseFloat(match.Price, 64)
        // Find High value
        if EthUsdCheck > EthUsdHigh { EthUsdHigh = EthUsdCheck }
        // Find Low value
        if EthUsdLow == 0 { EthUsdLow = EthUsdHigh }
        if EthUsdCheck < EthUsdLow { EthUsdLow = EthUsdCheck }
    // Fudge Typical Price
    // https://corporatefinanceinstitute.com/resources/knowledge/trading-investing/volume-weighted-adjusted-price-vwap/
    EthUsdTP = (EthUsdHigh + EthUsdLow + EthUsdCheck) / 3
    // Fudge Cumulative Volume by storing last 200 in array
    // Ignore the fact that many are zero size at first
    // Use Modulus to continue to reuse array after hitting 200 record mark
    if EthUsdCnt < 200 {
    EthUsdVol[EthUsdCnt],err = strconv.ParseFloat(match.Size, 64)
    } else  {
    EthUsdVol[EthUsdCnt % 200],err = strconv.ParseFloat(match.Size, 64)
    }

    for _, value := range EthUsdVol {
                EthUsdCum = EthUsdCum + value
        }

    EthUsdSize,err = strconv.ParseFloat(match.Size,64)
    fmt.Printf(" %f",EthUsdTP*EthUsdSize/EthUsdCum)

    fmt.Println()
    EthUsdCnt += 1
    } // End of ETH-USD
        // Figure out for ETH-BTC
    if strings.Contains(match.ProductId,"ETH-BTC") {
        EthBtcCheck,err = strconv.ParseFloat(match.Price, 64)
        // Find High value
        if EthBtcCheck > EthBtcHigh { EthBtcHigh = EthBtcCheck }
        // Find Low value
        if EthBtcLow == 0 { EthBtcLow = EthBtcHigh }
        if EthBtcCheck < EthBtcLow { EthBtcLow = EthBtcCheck }
    // Fudge Typical Price
    // https://corporatefinanceinstitute.com/resources/knowledge/trading-investing/volume-weighted-adjusted-price-vwap/
    EthBtcTP = (EthBtcHigh + EthBtcLow + EthBtcCheck) / 3
    // Fudge Cumulative Volume by storing last 200 in array
    // Ignore the fact that many are zero size at first
    // Use Modulus to continue to reuse array after hitting 200 record mark
    if EthBtcCnt < 200 {
    EthBtcVol[EthBtcCnt],err = strconv.ParseFloat(match.Size, 64)
    } else  {
    EthBtcVol[EthBtcCnt % 200],err = strconv.ParseFloat(match.Size, 64)
    }

    for _, value := range EthBtcVol {
                EthBtcCum = EthBtcCum + value
        }

    EthBtcSize,err = strconv.ParseFloat(match.Size,64)
    fmt.Printf(" %f",EthBtcTP*EthBtcSize/EthBtcCum)

    fmt.Println()
    EthBtcCnt += 1
    } // End of ETH-BTC
    }
    }

}

