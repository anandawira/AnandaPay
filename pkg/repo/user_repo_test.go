package repo

import (
	"context"
	"log"
	"testing"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		const email string = "duplicate@gmail.com"
		user1 := model.User{
			FullName:       "User1",
			Email:          email,
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}
		user2 := model.User{
			FullName:       "User2",
			Email:          email,
			HashedPassword: "hashedPassword2",
			IsVerified:     false,
		}

		err := ts.repo.Insert(context.TODO(), user1.FullName, user1.Email, user1.HashedPassword, user1.IsVerified)
		require.NoError(t, err)

		err = ts.repo.Insert(context.TODO(), user2.FullName, user2.Email, user2.HashedPassword, user2.IsVerified)
		assert.Error(t, err)
	})
}

func (ts *UserRepoTestSuite) TestGetOne() {
	ts.T().Run("It should return user and error nil if record found", func(t *testing.T) {
		user := model.User{
			FullName:       "User1",
			Email:          "email1@gmail.com",
			HashedPassword: "hashedPassword1",
			IsVerified:     false,
		}

		err := ts.repo.Insert(context.TODO(), user.FullName, user.Email, user.HashedPassword, user.IsVerified)
		require.NoError(t, err)

		result, err := ts.repo.GetByEmail(context.TODO(), user.Email)
		require.NoError(t, err)
		assert.Equal(t, user.Email, result.Email)
		assert.Equal(t, user.HashedPassword, result.HashedPassword)
	})

	ts.T().Run("It should return error if record not found", func(t *testing.T) {
		_, err := ts.repo.GetByEmail(context.TODO(), "noemail@gmail.com")
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
