package task3

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func (stu *Student) addStu(s Student) error {
	initDb()
	tx := db.Create(&s)
	if tx.Error != nil {
		fmt.Println("插入失败")
		return tx.Error
	}
	if tx.RowsAffected > 0 {
		fmt.Println("添加成功")

	}
	return nil
}

/**
查询 students 表中所有年龄大于 18 岁的学生信息
*/

func (stu *Student) findStuByAge(age int) (st *Student, e error) {
	initDb()
	var stt Student = Student{}
	tx := db.Where("age > ?", age).Find(&stt)
	if tx.RowsAffected > 0 {
		fmt.Println("查找数据成功:", stt.Age, stt.Name)
		return st, nil
	}
	if tx.Error != nil {
		fmt.Println("查找数据失败")
		return nil, tx.Error
	}
	return &stt, nil
}

func (stu Student) udateStudent() (n int, e error) {
	initDb()
	var s Student = Student{Grade: "四年级"}
	tx := db.Model(stu).Where("name = ?", "张三").Updates(&s)
	if tx.RowsAffected > 0 {
		fmt.Println("更新数据成功:", s.Age, s.Name)
		return int(tx.RowsAffected), nil
	}
	return 0, tx.Error
}

func (stu Student) delStudent() (n int, e error) {
	initDb()
	var s Student = Student{Grade: "四年级"}
	tx := db.Where("age < ?", 15).Unscoped().Delete(&s)
	if tx.RowsAffected > 0 {
		fmt.Println("删除数据成功:", s.Age, s.Name)
		return int(tx.RowsAffected), nil
	}
	return 0, tx.Error
}

/*
*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func transfer(toAccountId int32, fromAccountId int32, amount decimal.Decimal) error {
	initDb()
	tx := db.Begin()
	defer tx.Rollback() // 确保异常时自动回滚

	// 事务内查询付款账户
	var fromAcc Accounts
	if err := tx.Where("id = ?", fromAccountId).First(&fromAcc).Error; err != nil {
		return err
	}
	if fromAcc.Balance.LessThan(amount) {
		return errors.New("余额不足")
	}

	// 原子更新付款账户余额
	result := tx.Model(&Accounts{}).Where("id = ?", fromAccountId).
		Update("balance", gorm.Expr("balance - ?", amount))
	if result.Error != nil {
		return result.Error
	}

	// 原子更新收款账户余额并验证存在性
	result = tx.Model(&Accounts{}).Where("id = ?", toAccountId).
		Update("balance", gorm.Expr("balance + ?", amount))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("收款账户不存在")
	}

	// 记录转账
	trans := Transactions{
		FromAccountId: fromAccountId,
		ToAccountId:   toAccountId,
		Balance:       amount,
	}
	if err := tx.Create(&trans).Error; err != nil {
		return err
	}

	// 显式提交并返回错误（提交失败时自动回滚）
	return tx.Commit().Error
}
