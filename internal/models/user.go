package models

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (user *User) Marshall(isPublic bool) *User {
	user.Password = ""
	if isPublic {
		user.Email = ""
		user.CreatedAt = ""
	}

	return user
}
