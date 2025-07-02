package task3

import "github.com/shopspring/decimal"

type Accounts struct {
	Id      int32           `gorm:"primaryKey"`
	Balance decimal.Decimal `gorm:"type:decimal(10,2)"` //映射到数据库 Decimal(10,2)
}

type Transactions struct {
	Id            int32           `gorm:"primaryKey"`
	FromAccountId int32           `gorm:"column:from_account_id"`
	ToAccountId   int32           `gorm:"column:to_account_id"`
	Balance       decimal.Decimal `gorm:"type:decimal(10,2)"`
}

func (Accounts) TableName() string {
	return "accounts"
}

func (Transactions) TableName() string {
	return "transactions"
}
