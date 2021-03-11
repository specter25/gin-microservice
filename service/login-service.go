package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "ujjwal",
		authorizedPassword: "agarwal",
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedPassword == password && service.authorizedUsername == username
}
