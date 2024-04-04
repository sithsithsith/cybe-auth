package user_service

import (
	"github.com/google/uuid"
	"github.com/sithsithsith/cybe-auth/core/lib/exceptions"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"user_id"`
	Handle   string `json:"user_name"`
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

var users = []User{
	{
		ID:       uuid.NewString(),
		Handle:   "iAmFirstUser",
		Email:    "firstUser@example.com",
		Password: "password123",
	},
	{
		ID:       uuid.NewString(),
		Handle:   "iAmSecondUser",
		Email:    "secondUser@example.com",
		Password: "password123",
	},
	{
		ID:       uuid.NewString(),
		Handle:   "iAmThirdUser",
		Email:    "thirdUser@example.com",
		Password: "password123",
	},
}

func GetUsersList() []User {
	return users
}

func NewUser() *User {
	return &User{}
}

func CreateUser(u User) (User, error) {
	users = append(users, u)
	if err := len(users) == 4; err {
		return User{}, exceptions.NewUserException("Failed to create User")
	}
	return u, nil
}

func (u *User) HashPassword(cost int) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return exceptions.NewUserException("Password Hash Failed")
	}
	u.Password = string(bytes)
	return nil
}
