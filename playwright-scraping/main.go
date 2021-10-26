package main

import (
	"asset_secretary/common"
	"fmt"
	"strconv"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	var stockId string
	fmt.Println("Enter stock id: ")
	fmt.Scanln(&stockId)

	if len(stockId) == 0 {
		panic("Please provide stock id!")
	}

	baseUrl := "https://goodinfo.tw/StockInfo/StockDetail.asp?STOCK_ID="
	// Run PlayWright
	pw, err := playwright.Run()
	common.CheckError(err)
	
	// create browser object
	browser, err := pw.Chromium.Launch()
	common.CheckError(err)

	page, err := browser.NewPage()
	common.CheckError(err)

	// navigate to goodinfo for stock info
	page.Goto(fmt.Sprintf("%v%v", baseUrl, stockId))

	// document.querySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(1)")
	priceElement, err := page.QuerySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(1)")
	common.CheckError(err)
	price, err := priceElement.TextContent()
	common.CheckError(err)

	fmt.Println(fmt.Sprintf("price: %v", price))
	priceFloat, _ := strconv.ParseFloat(price, 32)
	fmt.Println(fmt.Sprintf("price after parse: %v", priceFloat))
	fmt.Println(fmt.Sprintf("price after parse and float(): %v", float32(priceFloat)))

	// document.querySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(4)")
	percentageElement, err := page.QuerySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(4)")
	common.CheckError(err)
	percentage, err := percentageElement.TextContent()
	common.CheckError(err)

	fmt.Println(fmt.Sprintf("percentage: %v", percentage))

}