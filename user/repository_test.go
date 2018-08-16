package user_test

import (
	"errors"
	user "golang-training/user"
	"log"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type RepositoryMock struct {
	DB *gorm.DB
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
	return user.User{}, errors.New("Error")
}

func (rm *RepositoryMock) CreateUser(user *user.User) error {
	return nil
}

func (rm *RepositoryMock) UpdateUser(user *user.User) error {
	return nil
}

func (rm *RepositoryMock) CreatePost(user *user.User, content string) error {
	return nil
}

func (rm *RepositoryMock) GetPosts(conditions map[string]interface{}) ([]user.Post, error) {
	return []user.Post{}, nil
}

func (rm *RepositoryMock) FindPost(id int) (user.Post, error) {
	return user.Post{}, nil
}

func (rm *RepositoryMock) UpdatePost(post *user.Post, content string) (user.Post, error) {
	return user.Post{}, nil
}

func (rm *RepositoryMock) DeletePost(post *user.Post) error {
	return nil
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
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at", "deleted_at", "fullname"}).AddRow(1, "", "", time.Time{}, time.Time{}, time.Time{}, "")
	mock.ExpectQuery("^SELECT (.+)").WillReturnRows(rows)
	_, err := h.Repository.FindUser(map[string]interface{}{"id": 1})
	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	// mock.ExpectQuery("^INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectExec("^INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	err := h.Repository.CreateUser(&user.User{
		Username: "test",
		Password: "test",
	})
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	mock.ExpectExec("^UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	err := h.Repository.UpdateUser(&user.User{
		FullName: "update",
	}, &user.User{
		Username: "a",
		Password: "a",
	})
	assert.NoError(t, err)
}

func TestCreatePost(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	mock.ExpectExec("^INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	post, err := h.Repository.CreatePost(&user.User{ID: 1}, "content")
	assert.NoError(t, err)
	assert.Equal(t, 1, post.IdUser)
	assert.Equal(t, "content", post.Content)
}

func TestGetPosts(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "id_user", "content", "created_at", "updated_at", "delete_at"}).AddRow(1, 1, "content", time.Time{}, time.Time{}, time.Time{})
	mock.ExpectQuery("^SELECT").WillReturnRows(rows)
	_, err := h.Repository.GetPosts(map[string]interface{}{})
	assert.NoError(t, err)
}

func TestFindPost(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "id_user", "content", "created_at", "updated_at", "delete_at"}).AddRow(1, 1, "content", time.Time{}, time.Time{}, time.Time{})
	mock.ExpectQuery("^SELECT").WillReturnRows(rows)
	_, err := h.Repository.FindPost(1)
	assert.NoError(t, err)
}

func TestUpdatePost(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	mock.ExpectExec("^UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	_, err := h.Repository.UpdatePost(&user.Post{}, "content")
	assert.NoError(t, err)
}

func TestDeletePost(t *testing.T) {
	mock, db := newDB()
	defer db.Close()
	h := user.Handler{
		Repository: &user.Repository{
			DB: db,
		},
	}
	mock.ExpectExec("^UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	err := h.Repository.DeletePost(&user.Post{})
	assert.NoError(t, err)
}
