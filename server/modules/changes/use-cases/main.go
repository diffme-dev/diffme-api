package use_cases

import (
	"diffme.dev/diffme-api/server/modules/changes"
)

type ChangeUseCases struct {
	changeRepo       domain.ChangeRepository
	searchChangeRepo domain.SearchChangeRepository
}

func NewChangeUseCase(changeRepo domain.ChangeRepository, searchChangeRepo domain.SearchChangeRepository) *ChangeUseCases {
	return &ChangeUseCases{
		changeRepo:       changeRepo,
		searchChangeRepo: searchChangeRepo,
	}
}
