package repo

import (
	"log"
	"testing"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/repo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WalletRepoTestSuite struct {
	suite.Suite

	DB   *gorm.DB
	repo domain.WalletRepository

	wallet1 domain.Wallet
	wallet2 domain.Wallet
}

func (ts *WalletRepoTestSuite) SetupSuite() {
	// Hardcore, later change to env variable
	dsn := "root:example@tcp(127.0.0.1:3306)/anandapay-test-wallet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	ts.DB = db
	ts.repo = NewWalletRepository(db)
	ts.DB.Migrator().DropTable(&domain.User{}, &domain.Wallet{})
	ts.DB.AutoMigrate(&domain.User{}, &domain.Wallet{}, &domain.Transaction{})

	userRepo := repo.NewUserRepository(db)
	err = userRepo.Insert(
		"fullname1",
		"email1@gmail.com",
		"hashedPassword1",
		false,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = userRepo.Insert(
		"fullname2",
		"email2@gmail.com",
		"hashedPassword2",
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	res := db.First(&ts.wallet1)
	if res.Error != nil {
		log.Fatal(err)
	}

	res = db.Last(&ts.wallet2)
	if res.Error != nil {
		log.Fatal(err)
	}
}

func (ts *WalletRepoTestSuite) TearDownSuite() {
	ts.DB.Migrator().DropTable(&domain.User{}, &domain.Wallet{})
	conn, err := ts.DB.DB()
	if err != nil {
		log.Fatal("Database not found")
	}
	conn.Close()
}

func (ts *WalletRepoTestSuite) TestGetBalance() {
	ts.T().Run("It should return balance and error nil on wallet found", func(t *testing.T) {
		balance, err := ts.repo.GetBalance(ts.wallet1.ID)
		require.NoError(t, err)
		assert.Equal(t, ts.wallet1.Balance, balance)
	})

	ts.T().Run("It should return error on wallet not found", func(t *testing.T) {
		_, err := ts.repo.GetBalance("invalid id")
		require.Error(t, err)
	})
}

func (ts *WalletRepoTestSuite) TestTopUp() {
	const TOPUP_AMOUNT = 5000000
	ts.T().Run("It should add balance and return error nil on wallet found", func(t *testing.T) {
		initialBalance, err := ts.repo.GetBalance(ts.wallet1.ID)
		require.NoError(t, err)
		expectedBalance := initialBalance + TOPUP_AMOUNT

		err = ts.repo.TopUp(uuid.NewString(), time.Now(), ts.wallet1.ID, "notes", TOPUP_AMOUNT)
		require.NoError(t, err)

		gotBalance, err := ts.repo.GetBalance(ts.wallet1.ID)
		require.NoError(t, err)
		assert.Equal(t, expectedBalance, gotBalance)
	})

	ts.T().Run("It should return error on wallet not found", func(t *testing.T) {
		err := ts.repo.TopUp(uuid.NewString(), time.Now(), "invalid id", "notes", TOPUP_AMOUNT)
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(WalletRepoTestSuite))
}
