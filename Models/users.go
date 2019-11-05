package Models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"judge/util"
	"net/url"
	"time"
)

type User struct {
	Id         uint
	Login      string
	Password   string
	Name       string
	LastAction time.Time
}

type Token struct {
	Id        uint
	UserId    uint
	User      User
	Token     string
	ExpiredAt time.Time
}

func NewUser(login string, password string, name string) (Token, error) {
	var existingUser User
	err := GetDB().Where(User{Login: login}).First(&existingUser).Error
	if err != gorm.ErrRecordNotFound && err == nil {
		return Token{}, errors.New(fmt.Sprintf("User with login %s already exist", login))
	}
	passwordHash, err := util.GetHash(password)
	user := User{Login: login, Password: passwordHash, Name: name}
	err = GetDB().Create(&user).Error
	if err != nil {
		return Token{}, err
	}
	user.Password = ""
	token, _ := Login(login, password)
	return token, nil
}

func Login(login string, password string) (Token, error) {
	var user User
	err := GetDB().Where(User{Login: login}).First(&user).Error
	if err != nil {
		return Token{}, errors.New("Incorrect login or password ")
	}
	if !util.CompareHashAndPassword(user.Password, password) {
		return Token{}, errors.New("Incorrect login or password ")
	}
	user.LastAction = time.Now().UTC()
	GetDB().Save(&user)
	var token Token
	err = GetDB().Where(Token{UserId: user.Id}).First(&token).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Token{}, err
	}
	if err == gorm.ErrRecordNotFound {
		token.Token, _ = util.GetHash(util.RandomString(16))
		token.Token = url.QueryEscape(token.Token)
		token.ExpiredAt = time.Now().Add(time.Minute * 15).UTC()
		token.UserId = user.Id
		token.User = user
		GetDB().Create(&token)
		return token, nil
	} else {
		token.Token, _ = util.GetHash(util.RandomString(16))
		token.Token = url.QueryEscape(token.Token)
		token.ExpiredAt = time.Now().Add(time.Minute * 15).UTC()
		GetDB().Save(&token)
		return token, nil
	}
}
