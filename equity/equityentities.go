package equity

import (
	"encoding/json"
	"fmt"
	"github/nsego/common"
)

type EquityGLStocks struct {
	Data []Data                 `json:"data"`
	Time common.CustomTimeStamp `json:"time"`
}

type Data struct {
	Symbol           string               `json:"symbol"`
	Series           string               `json:"series"`
	OpenPrice        common.CustomFloat64 `json:"openPrice"`
	HighPrice        common.CustomFloat64 `json:"highPrice"`
	LowPrice         common.CustomFloat64 `json:"lowPrice"`
	Ltp              common.CustomFloat64 `json:"ltp"`
	PreviousPrice    common.CustomFloat64 `json:"previousPrice"`
	NetPrice         common.CustomFloat64 `json:"netPrice"`
	TradedQuantity   common.CustomFloat64 `json:"tradedQuantity"`
	TurnoverInLakhs  common.CustomFloat64 `json:"turnoverInLakhs"`
	AnnouncementDate common.CustomTime    `json:"lastCorpAnnouncementDate"`
	Announcement     string               `json:"lastCorpAnnouncement"`
}

const (
	NSE_GAINER_URL = "https://www1.nseindia.com/live_market/dynaContent/live_analysis/gainers/niftyGainers1.json"
	NSE_LOSER_URL  = "https://www1.nseindia.com/live_market/dynaContent/live_analysis/losers/niftyLosers1.json"
)

func (per *EquityGLStocks) Init(url string) {
	body, err := common.NewNSERequest(url)
	if err != nil {
		fmt.Println(err)

	}
	if err := json.Unmarshal(body, per); err != nil {
		fmt.Println(err)
	}

}
