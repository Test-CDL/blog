package response

import (
	"chrombit/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		UserName:  data.UserName,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
