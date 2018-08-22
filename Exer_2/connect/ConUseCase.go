package connect

type ConnectUseCase struct {
	Repo ConnectRepositoryInterface
}

type ConnectUseCaseInterface interface {
	ConnectUC() error
}

func (uc *ConnectUseCase) ConnectUC() error {
	err := uc.Repo.ConnectRepo()
	if err != nil {
		return err
	}
	return nil
}

func NewConUC(repo *ConnectRepository) *ConnectUseCase {
	return &ConnectUseCase{
		Repo: repo,
	}
}
