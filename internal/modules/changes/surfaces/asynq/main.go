package asynq

import domain "diffme.dev/diffme-api/internal/modules/changes"

type ChangeAsynqSurface struct {
	changeUseCases domain.ChangeUseCases
}

func NewChangeAsnqSurface(changeUseCases domain.ChangeUseCases) ChangeAsynqSurface {
	return ChangeAsynqSurface{
		changeUseCases: changeUseCases,
	}
}
