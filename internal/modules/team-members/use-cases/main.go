package use_cases

import (
	domain "diffme.dev/diffme-api/internal/modules/team-members"
)

type TeamMemberUseCases struct {
	teamMemberRepo domain.TeamMemberRepository
}

func NewTeamMemberUseCases(teamMemberRepo domain.TeamMemberRepository) domain.TeamMemberUseCases {
	return &TeamMemberUseCases{
		teamMemberRepo: teamMemberRepo,
	}
}
