package user_test

import (
	user "golang-training/user"
	"log"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type RepositoryMock struct {
	DB         *gorm.DB
	Errors     error
	User       user.User
	Conditions map[string]interface{}
}

func newDB() (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}

	gormDB, gerr := gorm.Open("mysql", db)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}

	gormDB.LogMode(true)

	return mock, gormDB
}

func (rm *RepositoryMock) FindUser(conditions map[string]interface{}) (user.User, error) {
	return rm.User, rm.Errors
}

func (rm *RepositoryMock) CreateUser(user *user.User) error {
	return rm.Errors
}

func (rm *RepositoryMock) UpdateUser(updatedUser *user.User) error {
	return rm.Errors
}

func TestNewRepository(t *testing.T) {
	_, db := newDB()
	expected := &user.Repository{
		DB: db,
	}
	actual := user.NewRepository(db)
	assert.Equal(t, expected, actual)
}

func TestFindUser(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	repo := user.NewRepository(db)
	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at", "deleted_at", "fullname"}).AddRow(1, "", "", time.Time{}, time.Time{}, time.Time{}, "")
	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)
	_, err := repo.FindUser(map[string]interface{}{"id": 1})
	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	repo := user.NewRepository(db)
	// mock.ExpectQuery("^INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec("INSERT INTO `users` \\(`username`,`password`,`full_name`(.+)\\)(.+)").WithArgs("test", "test", "test", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.CreateUser(&user.User{
		Username: "test",
		Password: "test",
		FullName: "test",
	})
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	repo := user.NewRepository(db)
	mock.ExpectExec("^UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.UpdateUser(&user.User{
		ID:       1,
		FullName: "update",
	})
	assert.NoError(t, err)
}
