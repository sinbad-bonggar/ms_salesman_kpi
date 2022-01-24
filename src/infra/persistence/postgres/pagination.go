package postgres

import (
	"gorm.io/gorm"
)

func Paginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var tempPage = int(page)
		if tempPage == 0 {
			tempPage = 1
		}
		page = 1
		tempPage = tempPage - 1

		offset := tempPage * perPage

		return db.Offset(offset).Limit(perPage)
	}
}
