package _transferAmount

import (
	"errors"

	"gorm.io/gorm"
)

type Account struct {
	Id      int
	Balance int
}

type transactions struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        int
}

func Run(db *gorm.DB) {
	//db.AutoMigrate(&Account{})
	//db.AutoMigrate(&transactions{})

	//db.Create(&Account{Id: 1, Balance: 1000})
	//db.Create(&Account{Id: 2, Balance: 1000})

	transfer(&Account{Id: 1, Balance: 1000}, &Account{Id: 2, Balance: 1000}, 100, db)
}

func transfer(a *Account, b *Account, amount int, db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		// 先查询账户的实际信息
		var fromAccount Account
		tx.Debug().First(&fromAccount, a.Id)
		if fromAccount.Balance < amount {
			return errors.New("余额不足")
		}

		tx.Debug().Model(&Account{}).Where("id =?", (*a).Id).Update("Balance", (*a).Balance-amount)
		tx.Debug().Model(&Account{}).Where("id =?", (*b).Id).Update("Balance", (*b).Balance+amount)
		tx.Debug().Create(&transactions{FromAccountId: (*a).Id, ToAccountId: (*b).Id, Amount: amount})
		return nil
	})
}
