package rest

// GetUserRequest ...
type GetUserRequest struct {
	Email string `json:"email" bson:"email"`
}
