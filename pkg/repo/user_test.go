package repo

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/anandawira/anandapay/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func NewMockRepo() (model.UserRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Unexpected error when opening a stub database connection: %s", err)
	}

	gormMock, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db, SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unexpected error when initializing gorm: %s", err)
	}

	repo := NewUserRepository(gormMock)

	return repo, mock, db
}

func TestInsert(t *testing.T) {
	repo, mock, db := NewMockRepo()
	defer db.Close()

	param := model.User{
		FullName:       "test name",
		Email:          "test@email.com",
		HashedPassword: "hashedpassword",
		IsVerified:     true,
	}
	queryString := "INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`full_name`,`email`,`hashed_password`,`is_verified`) VALUES (?,?,?,?,?,?,?)"
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(queryString)).WithArgs(AnyTime{}, AnyTime{}, param.DeletedAt, param.FullName, param.Email, param.HashedPassword, param.IsVerified).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Insert(context.TODO(), param.FullName, param.Email, param.HashedPassword, param.IsVerified)
	if err != nil {
		t.Errorf("Unexpected error while inserting row: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
