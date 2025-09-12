package uimport

import (
	"unchained/server/rimport"
	"unchained/server/tools/logger"
)

type UsecaseImport struct {
	Usecase
}

func NewUsecaseImport(
	ri *rimport.RepositoryImports,
	log *logger.Logger,
) *UsecaseImport {
	return &UsecaseImport{
		Usecase: Usecase{
			// Payme: usecase.NewPayme(ri, log),
			// Kicker: usecase.NewKicker(log, ri),
		},
	}
}
