package uimport

import (
	"unchained/server/config"
	"unchained/server/gimport"
	"unchained/server/internal/usecase"
	"unchained/server/rimport"
	"unchained/server/tools/logger"
)

type Usecase struct {
	Auth *usecase.Auth
}

func NewUsecaseImport(
	ri *rimport.Repository,
	ge *gimport.Getaway,
	log *logger.Logger,
	conf *config.Config,
) *Usecase {
	return &Usecase{
		Auth: usecase.NewAuth(ri, ge, log, conf),
	}
}
