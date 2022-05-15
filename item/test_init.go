package item

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestUnit struct {
	Mock            sqlmock.Sqlmock
	IItemRepository itemRepository
}

func NewTestUnit() TestUnit {
	tu := TestUnit{}
	// bersifat inisialisasi
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	dbGorm, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql-mock",
		ServerVersion:             "1.0.0",
		DSN:                       "mysql-mock",
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	tu.Mock = mock
	iFaceItemRepo := NewItemRepository(dbGorm)
	tu.IItemRepository = *iFaceItemRepo
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `items`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "fullname"}).
				AddRow(1, "sosis bakar").AddRow(2, "jus mangga"))
	return tu
}
