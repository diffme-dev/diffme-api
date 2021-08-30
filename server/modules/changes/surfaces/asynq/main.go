package asynq

import (
	"diffme.dev/diffme-api/server/modules/changes"
)

type ChangeAsynqSurface struct {
	changeUseCases domain.ChangeUseCases
}

func NewChangeAsnqSurface(changeUseCases domain.ChangeUseCases) ChangeAsynqSurface {
	return ChangeAsynqSurface{
		changeUseCases: changeUseCases,
	}
}
