package item

import (
	"database/sql/driver"
	"order_kafe/category"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Save(item Item) (Item, error)
// FetchAll() ([]Item, error)
// FindById(id int) (Item, error)
// Update(item Item) (Item, error)
// Delete(item Item) (Item, error)

func TestDelete(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewItemRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	_, err := repo.Delete(Item{
		ID:          1,
		Name:        "sosis bakar",
		Description: "saus BBQ",
		Price:       10000,
		CategoryID:  1,
		Category:    category.Categorie{},
		ImageUrl:    "",
		IsAvailable: 0,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	})
	assert.NoError(t, err)
	assert.True(t, true)
}

// func TestUpdate(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 	},
// 	})
// 	repo := NewItemRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
// 		WithArgs("abc", 1).
// 		WillReturnResult(sqlmock.NewResult(0, 1))
// 	fMock.ExpectCommit()

// 	_, err := repo.Update(Item{
// 		ID:          0,
// 		Name:        "abc",
// 		Description: "",
// 		Price:       0,
// 		CategoryID:  0,
// 		Category:    category.Categorie{},
// 		ImageUrl:    "",
// 		IsAvailable: 0,
// 		CreatedAt:   time.Time{},
// 		UpdatedAt:   time.Time{},
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }

func TestFindById(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `items` WHERE id = ? ORDER BY `items`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "sosis bakar"))
	// result query GORM nya seperti apa
	user, err := tu.IItemRepository.FindById(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestFetchAll(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `items`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "sosis bakar").AddRow(2, "jus mangga"))
	// result query GORM nya seperti apa
	listUser, err := tu.IItemRepository.FetchAll()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(listUser)
}
