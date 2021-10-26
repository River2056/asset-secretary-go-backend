package main

import (
	"asset_secretary/common"
	"asset_secretary/model"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mxschmitt/playwright-go"
)

func main() {
	var data []model.DailyPrice = make([]model.DailyPrice, 0)
	db, err := sql.Open("sqlite3", "../../stocks.db")
	common.CheckError(err)

	rows, err := db.Query("SELECT * FROM DAILY_PRICE")
	common.CheckError(err)

	for rows.Next() {
		row := model.DailyPrice{}
		err = rows.Scan(&row.Id, &row.StockId, &row.StockName, &row.CurrentPrice, &row.Percentage, &row.ModifyDate)
		common.CheckError(err)
		data = append(data, row)
	}

	pw, err := playwright.Run()
	common.CheckError(err)
	browser, err := pw.Chromium.Launch()
	common.CheckError(err)
	page, err := browser.NewPage()
	common.CheckError(err)

	for _, stock := range data {
		page.Goto(fmt.Sprintf("%v%v", common.BaseUrl, stock.StockId))
		priceElement, err := page.QuerySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(1)")
		common.CheckError(err)
		price, err := priceElement.TextContent()
		common.CheckError(err)

		// document.querySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(4)")
		percentageElement, err := page.QuerySelector("body > table:nth-child(8) > tbody > tr > td:nth-child(3) > table > tbody > tr:nth-child(1) > td > table > tbody > tr:nth-child(1) > td:nth-child(1) > table > tbody > tr:nth-child(3) > td:nth-child(4)")
		common.CheckError(err)
		percentage, err := percentageElement.TextContent()
		common.CheckError(err)

		priceFloat, _ := strconv.ParseFloat(price, 32)
		stock.CurrentPrice = float32(priceFloat)
		stock.Percentage = percentage

		// update database with new data
		stmt, err := db.Prepare("UPDATE DAILY_PRICE SET CURRENT_PRICE = ?, PERCENTAGE = ?, MODIFY_DATE = ? WHERE ID = ?")
		common.CheckError(err)
		currentTime := time.Now()
		stmt.Exec(stock.CurrentPrice, stock.Percentage, common.GetTimeString(&currentTime), stock.Id)
	}

	fmt.Println(fmt.Sprintf("done updating!"))
}