package usecase

type Usecase struct {
	provider Provider
}

func NewUsecase(prv Provider) *Usecase {
	return &Usecase{
		provider: prv,
	}
}

