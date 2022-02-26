package usecase

import (
	"context"
	"errors"
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

func (m *userUsecase) Register(c context.Context, fullname, email, plainPassword string) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		log.Fatal("Password hashing error", err.Error())
	}

	err = m.userRepo.Insert(ctx, fullname, email, string(hashedPassword), false)
	if err != nil {
		return errors.New("Email already in use.")
	}

	return nil
}

func (m *userUsecase) Login(ctx context.Context, email string, plainPassword string) (token string, err error) {
	user, err := m.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Incorrect email or password.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(plainPassword))
	if err != nil {
		return "", errors.New("Incorrect email or password.")
	}

	// Hardcode, later change to env
	var secretKey string = "secret"

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	token, err = claims.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal("JWT token generation failed.", err.Error())
	}

	return token, nil
}
