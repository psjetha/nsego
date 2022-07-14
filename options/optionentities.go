package options

import (
	"encoding/json"
	"fmt"
	"github/nsego/common"
)

type OptionChain struct {
	Records  Record   `json:"records"`
	Filtered Filtered `json:"filtered"`
}

type Filtered struct {
	Data     []Data `json:"data"`
	CETotVol Volume `json:"CE"`
	PETotVol Volume `json:"PE"`
}

type Record struct {
	ExpiryDates        []common.CustomTime `json:"expiryDates"`
	Data               []Data              `json:"data"`
	Timestamp          string              `json:"timestamp"`
	UnderlyingValue    float64             `json:"underlyingValue"`
	StrikePrices       []float64           `json:"strikePrices"`
	Index              Index               `json:"index"`
	CurrentExpiryDates common.CustomTime   `json:"-"`
}
type Data struct {
	StrikePrice int               `json:"strikePrice"`
	ExpiryDate  common.CustomTime `json:"expiryDate"`
	PE          *Option           `json:"PE,omitempty"`
	CE          *Option           `json:"CE,omitempty"`
}

type Volume struct {
	TotOI  int `json:"totOI"`
	TotVol int `json:"totVol"`
}

type Option struct {
	StrikePrice           int               `json:"strikePrice"`
	ExpiryDate            common.CustomTime `json:"expiryDate"`
	Underlying            string            `json:"underlying"`
	Identifier            string            `json:"identifier"`
	OpenInterest          float64           `json:"openInterest"`
	ChangeinOpenInterest  float64           `json:"changeinOpenInterest"`
	PchangeinOpenInterest float64           `json:"pchangeinOpenInterest"`
	TotalTradedVolume     int               `json:"totalTradedVolume"`
	ImpliedVolatility     float64           `json:"impliedVolatility"`
	LastPrice             float64           `json:"lastPrice"`
	Change                float64           `json:"change"`
	PChange               float64           `json:"pChange"`
	TotalBuyQuantity      int               `json:"totalBuyQuantity"`
	TotalSellQuantity     int               `json:"totalSellQuantity"`
	BidQty                int               `json:"bidQty"`
	Bidprice              float64           `json:"bidprice"`
	AskQty                int               `json:"askQty"`
	AskPrice              float64           `json:"askPrice"`
	UnderlyingValue       float64           `json:"underlyingValue"`
}

type Index struct {
	Key           string  `json:"key"`
	Index         string  `json:"index"`
	IndexSymbol   string  `json:"indexSymbol"`
	Last          float64 `json:"last"`
	Variation     float64 `json:"variation"`
	PercentChange float64 `json:"percentChange"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	PreviousClose float64 `json:"previousClose"`
	YearHigh      float64 `json:"yearHigh"`
	YearLow       float64 `json:"yearLow"`
	Pe            string  `json:"pe"`
	Pb            string  `json:"pb"`
	Dy            string  `json:"dy"`
	Declines      string  `json:"declines"`
	Advances      string  `json:"advances"`
	Unchanged     string  `json:"unchanged"`
}

const (
	OPTION_CHAIN_URL = "https://www.nseindia.com/api/option-chain-indices?symbol="
)

func (oc *OptionChain) Init(symbol string) {
	url := OPTION_CHAIN_URL + symbol

	fmt.Println(url)

	body, err := common.NewNSERequest(url)
	if err != nil {
		fmt.Println(err)

	}

	if err := json.Unmarshal(body, oc); err != nil {
		fmt.Println(err)
	}
	oc.Records.SetCurrentExpriryDate()

}

func (r *Record) SetCurrentExpriryDate() {
	r.CurrentExpiryDates = r.ExpiryDates[0]
}
