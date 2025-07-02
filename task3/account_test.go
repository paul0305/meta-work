package task3

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestTransfer(t *testing.T) {
	err := transfer(1, 3, decimal.NewFromFloat(50))
	if err != nil {
		t.Error(err)
	} else {
		t.Log("转账成功")
	}
}
