package representation

import "github.com/Tracking-Detector/td-backend/go/td_common/model"

type UserDataRepresentation struct {
	ID    string     `json:"_id" bson:"_id"`
	Email string     `json:"email" bson:"email"`
	Role  model.Role `json:"role" bson:"role"`
}

func ConvertUserDataToUserDataRepresentation(user *model.UserData) *UserDataRepresentation {
	return &UserDataRepresentation{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

func ConvertUserDatasToUserDataRepresentations(users []*model.UserData) []*UserDataRepresentation {
	userRepresentations := make([]*UserDataRepresentation, len(users))
	for i, u := range users {
		userRepresentations[i] = ConvertUserDataToUserDataRepresentation(u)
	}
	return userRepresentations
}
