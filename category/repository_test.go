package category

import (
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// FetchAll() ([]Categorie, error)
// FindById(id int) (Categorie, error)
// Delete(categorie Categorie) (Categorie, error)

// func TestSave(t *testing.T) {
// 	tu := NewTestUnit()
// 	category := Categorie{
// 		ID:   1,
// 		Name: "jus",
// 	}
// 	// ekspektasi query yg dijalankan sama si lib GORM
// 	tu.Mock.ExpectBegin()
// 	tu.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`id`,`name`) VALUES (?,?)")).
// 		WithArgs(category.ID, category.Name).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	tu.Mock.ExpectCommit()
// 	// result query GORM nya seperti apa
// 	_, err := tu.ICategoryRepository.Save(category)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	t.Log("success insert")
// }

func TestFetchAll(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `categories`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "snack").AddRow(2, "jus"))
	// result query GORM nya seperti apa
	listUser, err := tu.ICategoryRepository.FetchAll()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(listUser)
}

func TestFindById(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `categories` WHERE id = ? ORDER BY `categories`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "snack"))
	// result query GORM nya seperti apa
	user, err := tu.ICategoryRepository.FindById(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestDelete(t *testing.T) {
	dbMock, fMock, _ := sqlmock.New()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbMock,
		SkipInitializeWithVersion: true,
	},
	})
	repo := NewCategoryRepository(db)
	defer dbMock.Close()

	fMock.ExpectBegin()
	fMock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	fMock.ExpectCommit()

	_, err := repo.Delete(Categorie{ID: 1, Name: "snack"})
	assert.NoError(t, err)
	assert.True(t, true)
}
