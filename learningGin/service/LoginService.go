package service

import "log"

type LoginService interface {
	Login(username string, password string) bool
}

type loginServiceImpl struct {
	authorizedUsername string
	authorizedPassword string
}

func (l *loginServiceImpl) Login(username string, password string) bool {
	log.Println("received params", username, password)
	log.Println("actual values", l.authorizedUsername, l.authorizedPassword)
	auth := l.authorizedUsername == username && l.authorizedPassword == password
	log.Println("login attempt", auth)
	return auth
}

func NewLoginService() LoginService {
	return &loginServiceImpl{
		authorizedPassword: "test",
		authorizedUsername: "mofe",
	}
}
