package UseCases

import (
	"diffme.dev/diffme-api/internal/modules/changes"
)

type ChangeUseCases struct {
	changeRepo domain.ChangeRepository
}

func NewChangeUseCase(changeRepo domain.ChangeRepository) domain.ChangeUseCases {
	return &ChangeUseCases{
		changeRepo: changeRepo,
	}
}
