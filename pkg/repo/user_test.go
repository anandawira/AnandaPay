package repo

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite

	DB   *gorm.DB
	repo model.UserRepository
}

func (ts *UserRepoTestSuite) SetupSuite() {
	// Hardcore, later change to env variable
	dsn := "root:example@tcp(127.0.0.1:3306)/anandapay-test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	ts.DB = db
	ts.repo = NewUserRepository(db)
}

func (ts *UserRepoTestSuite) SetupTest() {
	ts.DB.Migrator().DropTable(&model.User{})
	ts.DB.AutoMigrate(&model.User{})
}

func (ts *UserRepoTestSuite) TearDownSuite() {
	conn, err := ts.DB.DB()
	if err != nil {
		log.Fatal("Database not found")
	}
	conn.Close()
}

func (ts *UserRepoTestSuite) TestInsert() {
	ts.T().Run("It should insert to the database.", func(t *testing.T) {
		user := model.User{
			FullName:       "User1",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}

		err := ts.repo.Insert(context.TODO(), user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		assert.NoError(t, err)
	})

	ts.T().Run("It should not insert to the database if email already exist.", func(t *testing.T) {
		user := model.User{
			FullName:       "User2",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword2",
			IsVerified:     false,
		}
		err := ts.repo.Insert(context.TODO(), user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		fmt.Println(err)
		// Check error
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

// type AnyTime struct{}

// // Match satisfies sqlmock.Argument interface
// func (a AnyTime) Match(v driver.Value) bool {
// 	_, ok := v.(time.Time)
// 	return ok
// }

// func NewMockRepo() (model.UserRepository, sqlmock.Sqlmock, *sql.DB) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatalf("Unexpected error when opening a stub database connection: %s", err)
// 	}

// 	gormMock, err := gorm.Open(mysql.New(mysql.Config{
// 		Conn: db, SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})

// 	if err != nil {
// 		log.Fatalf("Unexpected error when initializing gorm: %s", err)
// 	}

// 	repo := NewUserRepository(gormMock)

// 	return repo, mock, db
// }

// func TestInsert(t *testing.T) {
// 	repo, mock, db := NewMockRepo()
// 	defer db.Close()

// 	param := model.User{
// 		FullName:       "test name",
// 		Email:          "test@email.com",
// 		HashedPassword: "hashedpassword",
// 		IsVerified:     true,
// 	}
// 	queryString := "INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`full_name`,`email`,`hashed_password`,`is_verified`) VALUES (?,?,?,?,?,?,?)"
// 	mock.ExpectBegin()
// 	mock.ExpectExec(regexp.QuoteMeta(queryString)).WithArgs(AnyTime{}, AnyTime{}, param.DeletedAt, param.FullName, param.Email, param.HashedPassword, param.IsVerified).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	err := repo.Insert(context.TODO(), param.FullName, param.Email, param.HashedPassword, param.IsVerified)
// 	if err != nil {
// 		t.Errorf("Unexpected error while inserting row: %s", err)
// 	}

// 	if err = mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }
