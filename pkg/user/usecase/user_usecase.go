package usecase

import (
	"log"
	"strconv"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(repo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{userRepo: repo, contextTimeout: timeout}
}

func (m *userUsecase) Register(fullname, email, plainPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		log.Fatal("Password hashing error", err.Error())
	}

	err = m.userRepo.Insert(fullname, email, string(hashedPassword), false)
	if err != nil {
		return domain.ErrEmailUsed
	}

	return nil
}

func (m *userUsecase) Login(email string, plainPassword string) (domain.User, domain.Wallet, string, error) {
	user, wallet, err := m.userRepo.GetByEmail(email)
	if err != nil {
		return user, wallet, "", domain.ErrWrongEmailPass
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(plainPassword))
	if err != nil {
		return user, wallet, "", domain.ErrWrongEmailPass
	}

	// Hardcode, later change to env
	var secretKey string = "secret"

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.CustomJwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
		WalletID: wallet.ID,
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal("JWT token generation failed.", err.Error())
	}

	return user, wallet, token, nil
}
