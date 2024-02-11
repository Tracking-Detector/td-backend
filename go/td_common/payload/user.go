package payload

import "github.com/Tracking-Detector/td-backend/go/td_common/model"

type CreateUserData struct {
	Email string     `bson:"email"`
	Role  model.Role `bson:"role"`
}
