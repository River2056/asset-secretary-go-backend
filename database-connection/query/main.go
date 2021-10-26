package main

import (
	"asset_secretary/common"
	"database/sql"
	"fmt"

	"asset_secretary/model"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var data []model.DailyPrice
	db, err := sql.Open("sqlite3", "../stocks.db")
	common.CheckError(err)

	rows, err := db.Query("SELECT * FROM DAILY_PRICE")
	common.CheckError(err)

	for rows.Next() {
		row := model.DailyPrice{}
		err := rows.Scan(&row.Id, &row.StockId, &row.StockName, &row.CurrentPrice, &row.Percentage, &row.ModifyDate)
		common.CheckError(err)
		fmt.Println(row)
		data = append(data, row)
	}

	fmt.Println(fmt.Sprintf("price data: %v", data))
}