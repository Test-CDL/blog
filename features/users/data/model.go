package data

import (
	"chrombit/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique" json:"user_name"`
	Password string `json:"password"`
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		UserName:  data.UserName,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core users.Core) User {
	return User{
		UserName: core.UserName,
		Password: core.Password,
	}
}
