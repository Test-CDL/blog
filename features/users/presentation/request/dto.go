package request

import "chrombit/features/users"

type User struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

func ToCore(req User) users.Core {
	return users.Core{
		UserName: req.UserName,
		Password: req.Password,
	}
}
