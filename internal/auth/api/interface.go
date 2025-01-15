package api

type Usecase interface {
	Login(string) (string, error)
	Register(string) error
	GenerateJWT(string) (string, error)
}