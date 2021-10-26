package model

type DailyPrice struct {
	Id int
	StockId string
	StockName string
	CurrentPrice float32
	Percentage string
	ModifyDate string
}