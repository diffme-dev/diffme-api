package use_cases

import (
	domain "diffme.dev/diffme-api/server/modules/team-members"
)

func (u *TeamMemberUseCases) InviteTeamMember(params domain.TeamMember) (*domain.TeamMember, error) {

	// TODO:... add other stuff..
	return u.teamMemberRepo.Create(params)
}
