package user

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// func TestInsertUser(t *testing.T) {
// 	tu := NewTestUnit()
// 	user := User{
// 		ID:        1,
// 		Email:     "abida@mail.com",
// 		Password:  "5678",
// 		Fullname:  "abida",
// 		Role:      "user",
// 		Avatar:    "avatar",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	// ekspektasi query yg dijalankan sama si lib GORM
// 	tu.Mock.ExpectBegin()
// 	tu.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`,`fullname`,`role`,`id`,`avatar`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
// 		WithArgs(user.Email, user.Password, user.Fullname, user.Role, user.ID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	tu.Mock.ExpectCommit()
// 	// result query GORM nya seperti apa
// 	_, err := tu.IUserRepository.Save(user)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	t.Log("success insert")
// }

func TestFindByUserEmail(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "fullname", "email", "password", "role"}).
				AddRow(1, "fahmi hadi", "fahmi@gmail.com", "password", "admin"))
	// result query GORM nya seperti apa
	user, err := tu.IUserRepository.FindByEmail("fahmi@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

func TestFindById(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "fullname", "email", "password", "role"}).
				AddRow(1, "fahmi hadi", "fahmi@gmail.com", "password", "admin"))
	// result query GORM nya seperti apa
	user, err := tu.IUserRepository.FindById(1)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(user)
}

// func TestUpdate(t *testing.T) {
// 	dbMock, fMock, _ := sqlmock.New()
// 	db, _ := gorm.Open(mysql.New(mysql.Config{
// 		DriverName:                "mysql-mock",
// 		ServerVersion:             "1.0.0",
// 		DSN:                       "mysql-mock",
// 		Conn:                      dbMock,
// 		SkipInitializeWithVersion: true,
// 		DefaultStringSize:         0,
// 		DefaultDatetimePrecision:  nil,
// 		DisableDatetimePrecision:  false,
// 		DontSupportRenameIndex:    false,
// 		DontSupportRenameColumn:   false,
// 		DontSupportForShareClause: false,
// 	}), &gorm.Config{})
// 	repo := NewUserRepository(db)
// 	defer dbMock.Close()

// 	fMock.ExpectBegin()
// 	fMock.ExpectExec(regexp.QuoteMeta("UPDATE")).
// 		WithArgs("abc", 1).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	fMock.ExpectCommit()

// 	_, err := repo.Update(User{
// 		Fullname: "abc",
// 	})
// 	assert.NoError(t, err)
// 	assert.True(t, true)
// }
