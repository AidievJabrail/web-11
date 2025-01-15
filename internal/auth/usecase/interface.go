package usecase

type Provider interface {
	CheckUser(string, string) (string, error)
	CreateUser(string, string) error
	FoundUser(string) (int, error)
}