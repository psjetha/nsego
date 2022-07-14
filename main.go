package main

import (
	"fmt"
	"github/nsego/equity"
	"github/nsego/options"
	"math"
	"time"
)

func main() {
	optionsLst := new(options.OptionChain)
	optionsLst.Init("BANKNIFTY")

	eqNiftygainer := new(equity.EquityGLStocks)
	eqNiftygainer.Init(equity.NSE_GAINER_URL)
	eqNiftylooser := new(equity.EquityGLStocks)
	eqNiftylooser.Init(equity.NSE_LOSER_URL)
	fmt.Println(eqNiftygainer)
	fmt.Println(eqNiftylooser)

	for {
		time.Sleep(time.Second * 3)
		GetOptionChain(*optionsLst)

	}

}

func nearest_option(ind float64, step int) int {
	return int((math.Ceil(ind/float64(step)) * float64(step)))
}

func GetOptionChain(optionsLst options.OptionChain) {
	cemax := new(options.Option)
	pemax := new(options.Option)

	num := 50
	step := 10
	nearest := nearest_option(optionsLst.Records.Index.Last, num)

	fmt.Println("Nearest :", nearest)
	var optionChainList []options.Data

	for _, data := range optionsLst.Records.Data {

		if data.ExpiryDate.String() == optionsLst.Records.CurrentExpiryDates.String() {
			if (nearest-(step*num)) <= data.StrikePrice && data.StrikePrice <= (nearest+(step*num)) {

				optionChainList = append(optionChainList, data)
			}
			if data.CE != nil {
				if data.CE.OpenInterest > cemax.OpenInterest {
					cemax = data.CE
				}

			}
			if data.PE != nil {
				if data.PE.OpenInterest > pemax.OpenInterest {
					pemax = data.PE
				}
			}

		}
	}

	fmt.Println("CE Stricke Price Support", cemax, cemax.OpenInterest)
	fmt.Println("PE Stricke Price Resistance", pemax, pemax.OpenInterest)
	for _, item := range optionChainList {
		fmt.Println("CE :", item.StrikePrice, item.CE.OpenInterest, item.CE.TotalBuyQuantity, item.CE.TotalSellQuantity, item.CE.TotalTradedVolume)
		fmt.Println("PE :", item.StrikePrice, item.PE.OpenInterest, item.PE.TotalBuyQuantity, item.PE.TotalSellQuantity, item.PE.TotalTradedVolume)

	}

}
