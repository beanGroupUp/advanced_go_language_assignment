package main

import (
	"awesomeProject3/databaseDriver/transaction/model"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/**
fmt.Errorf 的使用：

go
return fmt.Errorf("failed to begin transaction: %v", err)
这里创建了一个错误对象并返回给调用者，让调用者决定如何处理。

log.Fatal 的使用：

go
log.Fatal("Failed to open database:", err)
这里在程序初始化阶段遇到无法连接的致命错误，直接终止程序。

总结区别
特性	fmt.Errorf	log.Fatal
返回值	返回 error 对象	不返回任何值（程序终止）
程序行为	继续执行	立即终止
defer 语句	正常执行	不会执行
使用场景	函数内部错误处理	程序初始化或不可恢复错误

*/

// sql写法 使用 db.Exec() 方法执行DDL（数据定义语言）
//func transaction(db *sql.DB, fromAccountID int, toAccountID int, amount float64) error {
//	//开始事务
//	tx, err := db.Begin()
//	if err != nil {
//		return fmt.Errorf("account not found:%v", err)
//	}
//
//	//检查转出账户余额是否足够
//	//FOR UPDATE：对查询到的行加锁（行级锁），防止其他事务修改或删除该行，直到当前事务提交或回滚。这是处理并发更新的关键操作（例如转账时锁定账户）。
//	var balance float64
//	//将查询结果中的列值赋给 balance 变量。
//	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = ? for update", fromAccountID).Scan(&balance)
//
//	if err != nil {
//		//异常回滚
//		err := tx.Rollback()
//		if err != nil {
//			return err
//		}
//		return fmt.Errorf("account not found:%v", err)
//	}
//	//检查账户余额是否足够
//	if balance < amount {
//		//如果余额不足则回滚
//		err := tx.Rollback()
//		if err != nil {
//			return err
//		}
//		return fmt.Errorf("insufficient balance in account %d; money:%v", fromAccountID, balance)
//	}
//
//	//如果余额足够，则从账户中扣除金额
//	_, err = tx.Exec("Update accounts SET balance = balance - ? WHERE id = ?", amount, fromAccountID)
//	if err != nil {
//		err := tx.Rollback()
//		if err != nil {
//			return err
//		}
//		//账户扣减失败
//		return fmt.Errorf("update account failed:%v", err)
//	}
//
//	//向转入账户增加金额
//	_, err = tx.Exec("update accounts set balance = balance + ? where id = ?", amount, toAccountID)
//	if err != nil {
//		err := tx.Rollback()
//		if err != nil {
//			return err
//		}
//		return fmt.Errorf("update account failed:%v", err)
//	}
//
//	//记录转账交易
//	tx.Exec(
//		"Insert into transactions (from_account_id, to_account_id,amount) values(?,?,?) ",
//		fromAccountID, toAccountID, amount,
//	)
//	if err != nil {
//		//更新失败回滚
//		err := tx.Rollback()
//		if err != nil {
//			return err
//		}
//		return fmt.Errorf("insert into transactions failed:%v", err)
//	}
//
//	//提交事务
//	if err := tx.Commit(); err != nil {
//		err := tx.Rollback()
//		if err != nil {
//			return err
//		}
//		return fmt.Errorf("commit failed:%v", err)
//	}
//	return nil
//}
//
//func main() {
//
//	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local")
//
//	if err != nil {
//		log.Fatal("Failed to open database", err)
//	}
//	defer db.Close()
//
//	err = db.Ping()
//	if err != nil {
//		log.Fatal("Failed to ping database", err)
//	}
//
//	//执行转账操作：从账户1向账户2转账100元
//	err = transaction(db, 1, 2, 100.0)
//	if err != nil {
//		log.Printf("transaction failed:%v", err)
//	} else {
//		log.Printf("transaction succeeded")
//	}
//
//}

//func init() {
//	var err error
//	var db *gorm.DB
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold:             time.Second,
//			LogLevel:                  logger.Info,
//			IgnoreRecordNotFoundError: true,
//			ParameterizedQueries:      false, //设置为 false，在 SQL 日志中显示实际参数
//			Colorful:                  true,
//		},
//	)
//	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
//
//	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("连接成功")
//	//创建Accounts 表
//	//db.AutoMigrate(&model.Accounts{})
//	//创建Transcations 表
//	db.AutoMigrate(&model.Transactions{})
//}

var db *gorm.DB

func initDB() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	dsn := "root:123456@tcp(127.0.0.1:3306)/orm_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	//建表
	//err = db.AutoMigrate(&model.Accounts{}, &model.Transactions{})
	if err != nil {
		panic(err)
	}
	return nil
}

//自动事务
//
//func TransferMoney(db *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) error {
//	//开始事务
//	return db.Transaction(func(tx *gorm.DB) error {
//		//检查转出账户是否存在并锁定：cite[1]
//		var fromAccount model.Accounts
//		if err := tx.First(&fromAccount, fromAccountId).Error; err != nil {
//			return err //账户不存在，自动回滚
//		}
//
//		//2.检查余额是否充足：cite[1]
//		if fromAccount.Balance < amount {
//			//余额不足，自动回滚
//			return fmt.Errorf("insufficient balance in account %d", fromAccountId)
//		}
//
//		//3.扣减转出账户余额：
//		if err := tx.Model(&fromAccount).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
//			return err //更新失败，自动回滚
//		}
//
//		//4.增加转入账户余额
//		var toAccount model.Accounts
//		if err := tx.First(&toAccount, toAccountId).Error; err != nil {
//			return err //账户不存在，自动回滚
//		}
//		if err := tx.Model(&toAccount).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
//			return err //更新失败，自动回滚
//		}
//
//		//5.记录交易信息
//		transaction := model.Transactions{
//			FromAccountId: fromAccountId,
//			ToAccountId:   toAccountId,
//			Amount:        amount,
//		}
//		if err := db.Create(&transaction).Error; err != nil {
//			return err //插入失败，自动回滚
//		}
//
//		//返回nil,自动提交事务
//		return nil
//	})
//}

// 手动事务
func TransferMoneyManual(db *gorm.DB, fromAccountId uint, toAccountId uint, amount float64) error {
	//开始事务
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	//出现错误时回表
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//1.检查出账户并锁定
	var fromAccount model.Accounts
	if err := tx.First(&fromAccount, fromAccountId).Error; err != nil {
		tx.Rollback()
		return err
	}

	//2.检查余额是否充足‘
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance in account %d", fromAccountId)
	}

	//3.扣减转出账户余额
	if err := tx.Model(&fromAccount).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	var toAccount model.Accounts
	if err := tx.First(&toAccount, toAccountId).Error; err != nil {
		tx.Rollback()
	}

	//4.增加转入账户余额
	if err := tx.Model(&toAccount).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	//5.记录交易信息
	transaction := model.Transactions{
		FromAccountId: fromAccountId,
		ToAccountId:   toAccountId,
		Amount:        amount,
	}

	if err := db.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	//提交事务
	return tx.Commit().Error
}

func main() {
	//初始化数据库连接池
	if err := initDB(); err != nil {
		panic("Failed to initialize database:" + err.Error())
	}

	//假设有两个账户
	fromAccountId := uint(1)
	toAccountId := uint(2)
	amount := 100.0

	//执行转账（使用推荐方案）
	err := TransferMoneyManual(db, fromAccountId, toAccountId, amount)
	if err != nil {
		fmt.Printf("TransferMoney err:%s\n", err.Error())
	} else {
		fmt.Printf("TransferMoney success:%f\n", amount)
	}
}
