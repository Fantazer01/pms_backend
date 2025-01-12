package postgres

import (
	"pms_backend/pms_api/internal/pkg/model"
	"strings"
)

func toUserShortsFromRepo(usersFromDb []userShort) []*model.UserShort {
	users := make([]*model.UserShort, len(usersFromDb))
	for i := range users {
		sb := strings.Builder{}
		sb.WriteString(usersFromDb[i].FirstName)
		if usersFromDb[i].MiddleName != "" {
			sb.WriteString(" ")
			sb.WriteString(usersFromDb[i].MiddleName)
		}
		sb.WriteString(" ")
		sb.WriteString(usersFromDb[i].LastName)
		users[i] = &model.UserShort{
			ID:       usersFromDb[i].ID,
			Username: usersFromDb[i].Username,
			IsAdmin:  usersFromDb[i].IsAdmin,
			FullName: sb.String(),
		}
	}
	return users
}

func toUserFromRepo(userFromDb user) *model.User {
	return &model.User{
		ID:         userFromDb.ID,
		Username:   userFromDb.Username,
		IsAdmin:    userFromDb.IsAdmin,
		FirstName:  userFromDb.FirstName,
		MiddleName: userFromDb.MiddleName,
		LastName:   userFromDb.LastName,
		Position:   userFromDb.Position,
		CreatedAt:  userFromDb.CreatedAt,
		UpdatedAt:  userFromDb.UpdatedAt,
	}
}

func toProjectShortsFromDb(projectsFromDb []project) []*model.Project {
	projects := make([]*model.Project, len(projectsFromDb))
	for i := range projects {
		projects[i] = &model.Project{
			ID:          projectsFromDb[i].ID,
			Name:        projectsFromDb[i].Name,
			IsActive:    projectsFromDb[i].IsActive,
			Description: projectsFromDb[i].Description,
			CreatedAt:   projectsFromDb[i].CreatedAt,
			UpdatedAt:   projectsFromDb[i].UpdatedAt,
		}
	}
	return projects
}
